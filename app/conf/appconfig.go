package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"
)

// AppConfig - the app's path/version config
type AppConfig struct {
	NavConfPath       string `json:"navConfPath"`
	RunningNavVersion string `json:"runningNavVersion"`
}

// InitAppConfig will create mock data for now
// to save into app-config.json for viper read
func InitAppConfig() error {

	// create mock config
	mockConfig := AppConfig{}

	// set mock navconf path (osx path for initial test)
	mockConfig.NavConfPath = "$HOME/Library/Application Support/NavCoin4/navcoin.conf"

	// set mock running nav version
	mockConfig.RunningNavVersion = "4.1.1"

	// save mock navconf path and running nav version to app-config.json
	saveAppConfig(mockConfig.NavConfPath, mockConfig.RunningNavVersion)

	return nil
}

// LoadAppConfig - load the config from app-config.json
// It will otherwise be created automatically
// depending on platform - win|osx|linux
func LoadAppConfig() error {

	// set app config name and paths
	viper.SetConfigName("app-config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./app")
	viper.AddConfigPath(".")

	// use viper to read in config
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	// extract config
	appconfig := parseAppConfig(AppConfig{})

	AppConf = appconfig

	return nil
}

// StartConfigManager fires off LoadAppConfig every 500ms
func StartConfigManager() {

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for range ticker.C {
			LoadAppConfig()
		}
	}()
}

// LoadRPCDetails tries to read the config file for the RPC server
// and extract the RPCuser and RPCpassword
func LoadRPCDetails(appconfig AppConfig) error {

	// get path to navcoin.conf from app-config.json
	var navconf = AppConf.NavConfPath

	// Read the RPC server config
	serverConfigFile, err := os.Open(navconf)
	if err != nil {
		return err
	}

	defer serverConfigFile.Close()
	content, err := ioutil.ReadAll(serverConfigFile)
	if err != nil {
		return err
	}

	// Extract the RPCuser
	rpcUserRegexp, err := regexp.Compile(`(?m)^\s*rpcuser=([^\s]+)`)
	if err != nil {
		return err
	}
	userSubmatches := rpcUserRegexp.FindSubmatch(content)
	if userSubmatches == nil {
		// No user found, nothing to do
		return errors.New("No RPC User set in the config")
	}

	// Extract the RPCpassword
	rpcPassRegexp, err := regexp.Compile(`(?m)^\s*rpcpassword=([^\s]+)`)
	if err != nil {
		return err
	}
	passSubmatches := rpcPassRegexp.FindSubmatch(content)
	if passSubmatches == nil {
		// No password found we will die
		return errors.New("No RPC Password set")
	}

	// save ther user and password into the user config
	NavConf.RpcUser = string(userSubmatches[1])
	NavConf.RpcPassword = string(passSubmatches[1])

	return nil

}

// parseAppConfig reads config settings from app-config.json
func parseAppConfig(appconf AppConfig) AppConfig {

	appconf.NavConfPath = viper.GetString("navconf")
	appconf.RunningNavVersion = viper.GetString("runningNavVersion")

	return appconf

}

// AppConfigData defines json model for app-config.json
type AppConfigData struct {
	NavConf           string `json:"navconf"`
	RunningNavVersion string `json:"runningNavVersion"`
}

// saveAppConfig saves navconf path and version to app-config.json
func saveAppConfig(confPath string, runningVersion string) error {

	// merge input path and version values with AppConfigData
	// use MarshalIndent to format json (prettyprint)
	jsonData, err := json.MarshalIndent(AppConfigData{
		NavConf:           confPath,
		RunningNavVersion: runningVersion,
	}, "", "\t")
	if err != nil {
		return err
	}

	// log out the pretty json
	fmt.Println(string(jsonData))

	// build path to app-config.json
	path := "app/app-config.json"
	log.Println("attempting to write new json data to " + path)

	// write jsonData to file path with WriteFile
	// set perm: fileMode to os0644 for overwrite
	_ = ioutil.WriteFile(path, jsonData, 0644)

	return nil

}

// TODO: StoreNavConfig will store rpcUser:rpcPassword in memory (TBD)
// func StoreNavConfig(navconfig NavConfig) {
//
// }
