package conf

import (
	"errors"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"regexp"
)

// AppConfig - the app's path/version config
type AppConfig struct {
	NavConfPath       string `json:"navConfPath"`
	RunningNavVersion string `json:"runningNavVersion"`
}

var AppConf AppConfig

// UserConfig - rpc credentials
type UserConfig struct {
	RpcUser     string `json:"rpcUser"`
	RpcPassword string `json:"rpcPassword"`
}

var UserConf UserConfig

// SetupViper: set config name & paths
func SetupViper() {
	viper.SetConfigName("app-config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./app")
	viper.AddConfigPath(".")
}

// LoadAppConfig - load the config from config-app.json
// It will otherwise be created automatically
// depending on platform - win|osx|linux
func LoadAppConfig() error {

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	appconfig := parseAppConfig(AppConfig{})

	AppConf = appconfig

	return nil
}

// loadNavConfig tries to read the config file for the RPC server
// and extract the RPC user and password from it.
func LoadRPCDetails(appconfig AppConfig) error {

	var navcoinconf = AppConf.NavConfPath

	println("navcoinconf", navcoinconf)

	// Read the RPC server config
	serverConfigFile, err := os.Open(navcoinconf)
	if err != nil {
		return err
	}

	defer serverConfigFile.Close()
	content, err := ioutil.ReadAll(serverConfigFile)
	if err != nil {
		return err
	}

	// Extract the rpcuser
	rpcUserRegexp, err := regexp.Compile(`(?m)^\s*rpcuser=([^\s]+)`)
	if err != nil {
		return err
	}
	userSubmatches := rpcUserRegexp.FindSubmatch(content)
	if userSubmatches == nil {
		// No user found, nothing to do
		return errors.New("No RPC User set in the config")
	}

	// Extract the rpcpass
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
	UserConf.RpcUser = string(userSubmatches[1])
	UserConf.RpcPassword = string(passSubmatches[1])

	return nil

}

// parseAppConfig reads config settings from config-app.json
func parseAppConfig(appconfig AppConfig) AppConfig {

	appconfig.NavConfPath = viper.GetString("navconf")
	appconfig.RunningNavVersion = viper.GetString("runningNavVersion")

	return appconfig

}
