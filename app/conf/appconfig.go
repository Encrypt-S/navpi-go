package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/spf13/viper"
)

// AppConfig defines a structure to store app config data
type AppConfig struct {
	NavConf           string   `json:"navconf"`
	RunningNavVersion string   `json:"runningNavVersion"`
	AllowedIps        []string `json:"allowedIps"`
	UIPassword        string   `json:"uiPassword"`
}

// StartConfigManager sets up the ticker loop to load app config
func StartConfigManager() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for range ticker.C {
			LoadAppConfig()
		}
	}()
}

// LoadAppConfig sets up viper, reads and parses app config
func LoadAppConfig() error {

	viper.SetConfigName("app-config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./app")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	appconfig := parseAppConfig(AppConfig{})

	AppConf = appconfig

	return nil
}

// MockAppConfig mocks out and saves the app config
func MockAppConfig() (AppConfig, error) {

	mockConfig := AppConfig{}
	mockConfig.NavConf = "$HOME/Library/Application Support/NavCoin4/navcoin.conf"
	mockConfig.RunningNavVersion = "4.1.1"
	//mockConfig.DetectedIp = "1.1.1.1.1"

	AppConf = mockConfig

	err := SaveAppConfig()
	if err != nil {
		log.Println("Unable to save mocked app config")
		log.Println("err", err)
	}

	return mockConfig, nil
}

// parseAppConfig takes AppConfig, parses and returns appconf
func parseAppConfig(appconf AppConfig) AppConfig {

	appconf.NavConf = viper.GetString("navconf")
	appconf.RunningNavVersion = viper.GetString("runningNavVersion")
	appconf.AllowedIps = viper.GetStringSlice("allowedIps")
	appconf.UIPassword = viper.GetString("uiPassword")

	return appconf

}

// SaveAppConfig formats/indents json and saves to app-config.json
func SaveAppConfig() error {

	jsonData, err := json.MarshalIndent(AppConfig{
		NavConf:           AppConf.NavConf,
		RunningNavVersion: AppConf.RunningNavVersion,
		AllowedIps:        AppConf.AllowedIps,
		UIPassword:        AppConf.UIPassword,
	}, "", "\t")
	if err != nil {
		return err
	}

	fmt.Println(string(jsonData))

	path := "app/app-config.json"

	log.Println("attempting to write json data to " + path)

	err = ioutil.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil

}
