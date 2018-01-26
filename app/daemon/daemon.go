package daemon

import (
	"time"
	"github.com/NAVCoin/navpi-go/app/conf"
	"log"
	"os/exec"
	"path/filepath"
	"os"
	"runtime"
	"errors"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"io"
	"archive/zip"
)

const (
	WindowsDaemonName string = "navcoind.exe"
)


type OSInfo struct {
	DaemonName string
	OS string
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


func DownloadAndStart(serverConfig *conf.ServerConfig, userConfig *conf.Config) (*exec.Cmd) {

	path, err := CheckForDaemon(serverConfig, userConfig)

	if(err != nil) {
		downloadDaemon(serverConfig, userConfig.RunningNavVersion)
	}else {
		return start(path)
	}

	return start(path)

}


func Stop(cmd *exec.Cmd) {

	if err := cmd.Process.Kill(); err != nil {
		log.Fatal("failed to kill: ", err)
	}
}


func CheckForDaemon (serverConfig *conf.ServerConfig, userConfig *conf.Config) (string, error) {

	log.Println("Checking daemon")

	// get the latest release info
	releaseVersion := userConfig.RunningNavVersion

	log.Println("Looking for version: " + releaseVersion)

	// get the apps current path
	path, err := getCurrentPath()
	if err != nil {
		return "", err
	}

	// build the path
	path += "/lib/navcoin-" + releaseVersion + "/bin/"+ getOSInfo().DaemonName
	log.Println("Looking for Navcoin Daemon at " + path )

	// check the daemon exists
	if !exists(path) {
		log.Println("No Daemon found for version: " + releaseVersion)
		return "", errors.New("No Daemon found for version: " + releaseVersion)
	}else {
		log.Println("Located Daemon for version: " + releaseVersion)
	}

	return path, nil

}


func start(daemonPath string) (*exec.Cmd)  {
	
	log.Println("Booting deamon")
	cmd := exec.Command(daemonPath)
	cmd.Start()

	return cmd

}



// Get the current os info and the daemon name for that os
func getOSInfo () OSInfo {

	osInfo := OSInfo{}

	//const goosList = "android darwin dragonfly freebsd linux nacl netbsd openbsd plan9 solaris windows zos "
	//const goarchList = "386 amd64 amd64p32 arm armbe arm64 arm64be ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le ppc s390 s390x sparc sparc64 "

	switch runtime.GOARCH {

	case "amd64":

		switch runtime.GOOS {
		case "windows":
			osInfo.DaemonName = WindowsDaemonName
			osInfo.OS = "win64"
		}

	}

	return osInfo

}

// gets the current path of the go application
func getCurrentPath() (string, error)  {

	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	exPath := filepath.Dir(ex)
	return exPath, nil
}


// Exists reports whether the named file or directory exists.
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}


func downloadDaemon(serverConf *conf.ServerConfig, verson string) {

	releaseInfo, _ := getReleaseDataForVersion(serverConf.ReleaseAPI, verson)

	dlPath, dlName, _ := getDwLdInfoFromReleaseInfo(releaseInfo)

	downloadUnzip(dlPath, dlName)

}

func getReleaseDataForVersion(releaseAPI string, version string) (GitHubReleaseData, error)  {

	log.Println("Atempting to get release data")
	releases, err := gitHubReleaseInfo(releaseAPI)

	var e GitHubReleaseData = GitHubReleaseData{}
	for _, elem := range releases {
		if elem.TagName == version {
			log.Println("Release data found for version: " + version)
			e = elem.GitHubReleaseData
		}
	}

	return e, err

}


func gitHubReleaseInfo(releaseAPI string) (GitHubReleases, error) {

	log.Println("Retreving NAVCoin Github releases info from: " + releaseAPI)

	response, err := http.Get(releaseAPI)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		return  GitHubReleases{}, err
	}

	// read the data out to json
	data, _ := ioutil.ReadAll(response.Body)
	c := GitHubReleases{}
	jsonErr := json.Unmarshal(data, &c)

	if jsonErr != nil {
		return  GitHubReleases{}, jsonErr
		log.Fatal(jsonErr)
	}

	return c, nil
}


func getDwLdInfoFromReleaseInfo(gitHubReleaseData GitHubReleaseData) (string, string, error) {

	log.Println("Getting dowload path for OS from release assest data")

	releaseInfo := gitHubReleaseData

	downloadPath := ""
	downloadName := ""

	for e := range releaseInfo.Assets {

		asset := releaseInfo.Assets[e]

		if strings.Contains(asset.Name, getOSInfo().OS) {
			// extra windows os check as we only want the zip
			if strings.Contains(asset.Name, "win") {
				//println(filepath.Ext(asset.Name))
				if filepath.Ext(asset.Name) == ".zip" {
					downloadPath = releaseInfo.Assets[e].BrowserDownloadURL
					downloadName = releaseInfo.Assets[e].Name
				}
			} else {
				downloadPath = releaseInfo.Assets[e].BrowserDownloadURL
				downloadName = releaseInfo.Assets[e].Name
			}
		}
	}

	log.Println("Download path found: " + downloadPath)

	return downloadPath, downloadName, nil

}


func downloadUnzip(assetPath string, assetName string) error {

	path, err := getCurrentPath()

	downloadLocation := path+ "/" + assetName

	log.Println("Downloading", assetPath, "to", downloadLocation)

	download(assetPath, downloadLocation)

	if filepath.Ext(assetName) == ".zip" {
		unzip(downloadLocation, path + "/lib")
	}

	if err != nil {
		return err
	}

	return nil

}


func download(url string, fileName string) {

	log.Println("Downloading daemon - this could take a few mins :)")

	output, err := os.Create(fileName)

	response, err := http.Get(url)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return
	}

	log.Println(n, "bytes downloaded")

}


func unzip(src, dest string) error {

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())
		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath,string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, f.Mode())
			if err != nil {
				log.Fatal(err)
				return err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}









/*

func NewDaemonAvailable(config *conf.ServerConfig) (bool, error) {

	log.Println("Checking daemon")

	// get the latest release info
	releaseVersion, versionErr := getReleaseVersion(config)
	if versionErr != nil {
		return false, versionErr
	}

	log.Println("Latest version: " + releaseVersion)

	// get the apps current path
	path, err := getCurrentPath()
	if err != nil {
		return false, err
	}

	// append the deamon name

	path += "/lib/navcoin-" + releaseVersion + "/bin/"+ getOSInfo().DaemonName
	log.Println("Looking for Navcoin Daemon at " + path )


	// check the daemon exists
	if !exists(path) {
		log.Println("No Daemon found for version: " + releaseVersion)

		return true, nil
	}else {
		log.Println("Located Daemon for version: " + releaseVersion)
		return false, nil
	}

}



func getDaemonDownloadPath(version string) {

	

}





func getReleaseAssetInfo(config *conf.ServerConfig) (string, string, error) {

	releaseInfo, err := gitHubReleaseInfo(config)
	if err != nil {
		return "", "", err
	}

	downloadPath := ""
	downloadName := ""

	for e := range releaseInfo.Assets {

		asset := releaseInfo.Assets[e]

		if strings.Contains(asset.Name, getOSInfo().OS) {

			// extra windows os check as we only want the zip
			if strings.Contains(asset.Name, "win") {
				//println(filepath.Ext(asset.Name))
				if filepath.Ext(asset.Name) == ".zip" {
					downloadPath = releaseInfo.Assets[e].BrowserDownloadURL
					downloadName = releaseInfo.Assets[e].Name
				}
			}else {
				downloadPath = releaseInfo.Assets[e].BrowserDownloadURL
				downloadName = releaseInfo.Assets[e].Name
			}

		}
	}

	return downloadPath, downloadName, nil

}

func gitHubReleaseInfo(serverConfig *conf.ServerConfig) (GitHubReleaseData, error) {

	log.Println("Retreving NAVCoin Github release info from: " + serverConfig.LatestReleaseAPI)

	response, err := http.Get(serverConfig.LatestReleaseAPI)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		return  GitHubReleaseData{}, err
	}

	// read the data out to json
	data, _ := ioutil.ReadAll(response.Body)
	c := GitHubReleaseData{}
	jsonErr := json.Unmarshal(data, &c)

	if jsonErr != nil {
		return  GitHubReleaseData{}, jsonErr
		log.Fatal(jsonErr)
	}

	return c, nil
}


func getReleaseVersion(serverConfig *conf.ServerConfig) (string, error) {

	releaseInfo, err := gitHubReleaseInfo(serverConfig)
	if err != nil {
		return "", nil
	}

	return releaseInfo.TagName, nil
}








*/

