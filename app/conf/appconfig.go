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
	NavConf           string `json:"navconf"`
	RunningNavVersion string `json:"runningNavVersion"`
	DetectedIp        string `json:"detectedIp"`
}

// StartConfigManager fires off LoadAppConfig every 500ms
// this loop is essential for detecting change in config
func StartConfigManager() {

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for range ticker.C {
			LoadAppConfig()
		}
	}()
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

	// update AppConf var
	AppConf = appconfig

	return nil
}

// MockAppConfig will create mock data for now
// to save into app-config.json for viper read
func MockAppConfig() (AppConfig, error) {

	// create mock config
	mockConfig := AppConfig{}

	// set mock navconf path (osx path for initial test)
	mockConfig.NavConf = "$HOME/Library/Application Support/NavCoin4/navcoin.conf"

	// set mock running nav version
	mockConfig.RunningNavVersion = "4.1.1"

	// set mock ip address
	mockConfig.DetectedIp = "1.1.1.1.1"

	// update the AppConf var
	AppConf = mockConfig

	// save the mocked app config
	err := saveAppConfig()
	if err != nil {
		log.Println("Unable to save mocked app config")
		log.Println("err", err)
	}

	return mockConfig, nil
}

// parseAppConfig reads config settings from app-config.json
func parseAppConfig(appconf AppConfig) AppConfig {

	appconf.NavConf = viper.GetString("navconf")
	appconf.RunningNavVersion = viper.GetString("runningNavVersion")
	appconf.DetectedIp = viper.GetString("detectedIp")

	return appconf

}

// LoadRPCDetails tries to read the config file for the RPC server
// and extract the RPCuser and RPCpassword
func LoadRPCDetails(appconfig AppConfig) error {

	// get path to navcoin.conf from app-config.json
	var navconf = AppConf.NavConf

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

	// save the user and password into the user config
	NavConf.RpcUser = string(userSubmatches[1])
	NavConf.RpcPassword = string(passSubmatches[1])

	return nil

}

// SaveAppConfig saves navconf path and version to app-config.json
// will be called called only once after MockAppConfig
// next time the config loop runs :: we'll have a config
func saveAppConfig() error {

	// merge input path and version values with AppConfigData
	// use MarshalIndent to format json (prettyprint)
	jsonData, err := json.MarshalIndent(AppConfig{
		NavConf:           AppConf.NavConf,
		RunningNavVersion: AppConf.RunningNavVersion,
		DetectedIp:        AppConf.DetectedIp,
	}, "", "\t")
	if err != nil {
		return err
	}

	// log out the pretty json
	fmt.Println(string(jsonData))

	// build path to app-config.json
	path := "app/app-config.json"

	log.Println("attempting to write json data to " + path)

	// write jsonData to file path with WriteFile
	// set perm: fileMode to os0644 for overwrite
	err = ioutil.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil

}

// TODO: StoreNavConfig will store rpcUser:rpcPassword in memory (TBD)
// func StoreNavConfig(navconfig NavConfig) {
//
// }
