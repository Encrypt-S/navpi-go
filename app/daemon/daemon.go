package daemon

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/NAVCoin/navpi-go/app/conf"
	"github.com/NAVCoin/navpi-go/app/daemon/daemonrpc"
	"github.com/NAVCoin/navpi-go/app/fs"
)

const (
	WindowsDaemonName string = "navcoind.exe"
	DarwinDaemonName  string = "navcoind"
)

type OSInfo struct {
	DaemonName string
	OS         string
}

type GitHubReleases []struct {
	GitHubReleaseData
}

type GitHubReleaseData struct {
	URL             string `json:"url"`
	AssetsURL       string `json:"assets_url"`
	UploadURL       string `json:"upload_url"`
	HTMLURL         string `json:"html_url"`
	ID              int    `json:"id"`
	TagName         string `json:"tag_name"`
	TargetCommitish string `json:"target_commitish"`
	Name            string `json:"name"`
	Draft           bool   `json:"draft"`
	Author          struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	Prerelease  bool      `json:"prerelease"`
	CreatedAt   time.Time `json:"created_at"`
	PublishedAt time.Time `json:"published_at"`
	Assets      []struct {
		URL      string      `json:"url"`
		ID       int         `json:"id"`
		Name     string      `json:"name"`
		Label    interface{} `json:"label"`
		Uploader struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"uploader"`
		ContentType        string    `json:"content_type"`
		State              string    `json:"state"`
		Size               int       `json:"size"`
		DownloadCount      int       `json:"download_count"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		BrowserDownloadURL string    `json:"browser_download_url"`
	} `json:"assets"`
	TarballURL string `json:"tarball_url"`
	ZipballURL string `json:"zipball_url"`
	Body       string `json:"body"`
}

var runningDaemon *exec.Cmd
var minHeartbeat int64 = 500 // the lowest value the hb checker can be set to

// StartManager is a simple system that checks if
// the daemon is alive. If not it tries to start it
func StartManager() {

	// set the heartbeat interval but make sure it is not
	// less than the min heartbeat setting
	hbInterval := minHeartbeat
	if conf.ServerConf.DaemonHeartbeat > hbInterval {
		hbInterval = conf.ServerConf.DaemonHeartbeat
	}

	ticker := time.NewTicker(time.Duration(hbInterval) * time.Millisecond)
	go func() {
		for t := range ticker.C {

			log.Println(t)

			// check to see if the daemon is alive
			if isAlive() {
				log.Println("NAVCoin daemon is alive!")
			} else {

				log.Println("NAVCoin daemon is unresponsive...")

				if runningDaemon != nil {
					Stop(runningDaemon)
				}

				// start the daemon and download it if necessary
				cmd, err := DownloadAndStart(conf.ServerConf, conf.AppConf)

				if err != nil {
					log.Println(err)
				} else {
					runningDaemon = cmd
				}

			}

		}
	}()
}

// isAlive performs a simple rpc command to the Daemon
// returns false on error
func isAlive() bool {

	isLiving := true

	n := daemonrpc.RpcRequestData{}
	n.Method = "getblockcount"

	_, err := daemonrpc.RequestDaemon(n, conf.NavConf)

	if err != nil {
		isLiving = false
	}

	return isLiving

}

func DownloadAndStart(serverConfig conf.ServerConfig, appConfig conf.AppConfig) (*exec.Cmd, error) {

	if appConfig.RunningNavVersion == "" {
		return nil, errors.New("no nav version set in the user config")
	}

	path, err := CheckForDaemon(serverConfig, appConfig)

	if err != nil {
		downloadDaemon(serverConfig, appConfig.RunningNavVersion)
	} else {
		return start(path), nil
	}

	return start(path), nil

}

func Stop(cmd *exec.Cmd) {

	if err := cmd.Process.Kill(); err != nil {
		log.Fatal("failed to kill: ", err)
	}
}

func CheckForDaemon(serverConfig conf.ServerConfig, appConfig conf.AppConfig) (string, error) {

	// get the latest release info
	releaseVersion := appConfig.RunningNavVersion

	log.Println("Checking NAVCoin daemon for v" + releaseVersion)

	// get the apps current path
	path, err := fs.GetCurrentPath()
	if err != nil {
		return "", err
	}

	// build the path
	path += "/lib/navcoin-" + releaseVersion + "/bin/" + getOSInfo().DaemonName
	log.Println("Searching for NAVCoin daemon at " + path)

	// check the daemon exists
	if !fs.Exists(path) {
		log.Println("NAVCoin daemon not found for v" + releaseVersion)
		return "", errors.New("NAVCoin daemon found for v" + releaseVersion)
	} else {
		log.Println("NAVCoin daemon located for v" + releaseVersion)
	}

	return path, nil

}

func start(daemonPath string) *exec.Cmd {

	log.Println("Booting NAVCoin daemon")
	cmd := exec.Command(daemonPath)
	cmd.Start()

	return cmd

}

// Get the current os info and the Daemon name for that os
func getOSInfo() OSInfo {

	osInfo := OSInfo{}

	//const goosList = "android darwin dragonfly freebsd linux nacl netbsd openbsd plan9 solaris windows zos "
	//const goarchList = "386 amd64 amd64p32 arm armbe arm64 arm64be ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le ppc s390 s390x sparc sparc64"

	switch runtime.GOARCH {

	case "amd64":

		switch runtime.GOOS {

		case "windows":

			osInfo.DaemonName = WindowsDaemonName
			osInfo.OS = "win64"
			break

		case "darwin":

			osInfo.DaemonName = DarwinDaemonName
			osInfo.OS = "osx64"
			break
		}

		break
	}

	return osInfo

}

func downloadDaemon(serverConf conf.ServerConfig, version string) {

	releaseInfo, _ := getReleaseDataForVersion(serverConf.ReleaseAPI, version)

	dlPath, dlName, _ := getDownloadPathAndName(releaseInfo)

	fs.DownloadExtract(dlPath, dlName)

}

func getReleaseDataForVersion(releaseAPI string, version string) (GitHubReleaseData, error) {

	log.Println("Attempting to get release data for NAVCoin v" + version)

	releases, err := gitHubReleaseInfo(releaseAPI)

	var e GitHubReleaseData = GitHubReleaseData{}

	for _, elem := range releases {
		if elem.TagName == version {
			log.Println("Release data found for NAVCoin v" + version)
			e = elem.GitHubReleaseData
		}
	}

	return e, err

}

func gitHubReleaseInfo(releaseAPI string) (GitHubReleases, error) {
	log.Println("Retrieving NAVCoin Github release data from: " + releaseAPI)
	response, err := http.Get(releaseAPI)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		return GitHubReleases{}, err
	}

	// read the data out to json
	data, _ := ioutil.ReadAll(response.Body)
	c := GitHubReleases{}
	jsonErr := json.Unmarshal(data, &c)

	if jsonErr != nil {
		return GitHubReleases{}, jsonErr
		log.Fatal(jsonErr)
	}

	return c, nil
}

func getDownloadPathAndName(gitHubReleaseData GitHubReleaseData) (string, string, error) {

	log.Println("Getting download path/name for OS from release assest data")

	releaseInfo := gitHubReleaseData

	downloadPath := ""
	downloadName := ""

	for e := range releaseInfo.Assets {

		asset := releaseInfo.Assets[e]

		if strings.Contains(asset.Name, getOSInfo().OS) {
			// windows os check to provide .zip
			if strings.Contains(asset.Name, "win") {
				if filepath.Ext(asset.Name) == ".zip" {
					log.Println("win64 detected - preparing NAVCoin .zip download")
					downloadPath = releaseInfo.Assets[e].BrowserDownloadURL
					downloadName = releaseInfo.Assets[e].Name
				}
			}
			// osx64 check to provide gzip package :: tar.gz
			if strings.Contains(asset.Name, "osx64") {
				log.Println("osx64 detected - preparing NAVCoin tar.gz download")
				downloadPath = releaseInfo.Assets[e].BrowserDownloadURL
				downloadName = releaseInfo.Assets[e].Name
			} else {
				// TODO: more checks to be added for other systems
				// fall through to defaults :: fire-in-the-hole mode
				downloadPath = releaseInfo.Assets[e].BrowserDownloadURL
				downloadName = releaseInfo.Assets[e].Name
			}
		}
	}

	return downloadPath, downloadName, nil

}
