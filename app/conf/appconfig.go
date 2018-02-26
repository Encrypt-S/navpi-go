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

type AppConfig struct {
	NavConf           string `json:"navconf"`
	RunningNavVersion string `json:"runningNavVersion"`
	AllowedIps []string `json:"allowedIps"`
	UIPassword string `json:"uiPassword"`
}



func StartConfigManager() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for range ticker.C {
			LoadAppConfig()
		}
	}()
}

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

func parseAppConfig(appconf AppConfig) AppConfig {

	appconf.NavConf = viper.GetString("navconf")
	appconf.RunningNavVersion = viper.GetString("runningNavVersion")
	appconf.AllowedIps = viper.GetStringSlice("allowedIps")

	return appconf

}

func LoadRPCDetails(appconfig AppConfig) error {

	var navconf = AppConf.NavConf

	serverConfigFile, err := os.Open(navconf)
	if err != nil {
		return err
	}

	defer serverConfigFile.Close()
	content, err := ioutil.ReadAll(serverConfigFile)
	if err != nil {
		return err
	}

	rpcUserRegexp, err := regexp.Compile(`(?m)^\s*rpcuser=([^\s]+)`)
	if err != nil {
		return err
	}

	userSubmatches := rpcUserRegexp.FindSubmatch(content)
	if userSubmatches == nil {
		return errors.New("No RPC User set in the config")
	}

	rpcPassRegexp, err := regexp.Compile(`(?m)^\s*rpcpassword=([^\s]+)`)
	if err != nil {
		return err
	}

	passSubmatches := rpcPassRegexp.FindSubmatch(content)
	if passSubmatches == nil {
		return errors.New("No RPC Password set")
	}

	NavConf.RpcUser = string(userSubmatches[1])
	NavConf.RpcPassword = string(passSubmatches[1])

	return nil

}

func SaveAppConfig() error {

	jsonData, err := json.MarshalIndent(AppConfig{
		NavConf:           AppConf.NavConf,
		RunningNavVersion: AppConf.RunningNavVersion,
		AllowedIps:        AppConf.AllowedIps,
		UIPassword:			AppConf.UIPassword,
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
