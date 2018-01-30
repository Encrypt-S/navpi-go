package conf

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"regexp"
	"errors"
	"time"
)

// UserConfig the application's configuration
type UserConfig struct {
	NavConfPath       	string `json:"navConfPath"`
	RunningNavVersion 	string `json:"runningNavVersion"`
	RpcUser 				string `json:"rpcUser"`
	RpcPassword 			string `json:"rpcPassword"`
}

func StartConfigManager()  {

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for range ticker.C {
			LoadUserConfig()
		}
	}()
}


// LoadUserConfig loads the config from a file
func LoadUserConfig() (error)  {


	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./app")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		return err
	}

	// load the go server config
	config := parseConfig(UserConfig{})


	//load the navcoin daemon config
	//err = loadNavConfig(config)

	if err != nil {
		return err
	}


	// store the user config
	UserConf = config

	return nil
}

// loadNavConfig tries to read the config file for the RPC server
// and extract the RPC user and password from it.
func LoadRPCDetails (config *UserConfig) (string, string, error) {

	var configfile = config.NavConfPath

	// Read the RPC server config
	serverConfigFile, err := os.Open(configfile)
	if err != nil {
		return "", "", err
	}

	defer serverConfigFile.Close()
	content, err := ioutil.ReadAll(serverConfigFile)
	if err != nil {
		return "", "", err
	}

	// Extract the rpcuser
	rpcUserRegexp, err := regexp.Compile(`(?m)^\s*rpcuser=([^\s]+)`)
	if err != nil  {
		return "", "", err
	}
	userSubmatches := rpcUserRegexp.FindSubmatch(content)
	if userSubmatches == nil {
		// No user found, nothing to do
		return "", "", errors.New("No RPC User set in the config")
	}

	// Extract the rpcpass
	rpcPassRegexp, err := regexp.Compile(`(?m)^\s*rpcpassword=([^\s]+)`)
	if err != nil {
		return "", "", err
	}
	passSubmatches := rpcPassRegexp.FindSubmatch(content)
	if passSubmatches == nil {
		// No password found we will die
		return "", "", errors.New("No RPC Password set")
	}


	// save ther user and password into the app level config
	config.RpcUser = string(userSubmatches[1])
	config.RpcPassword= string(passSubmatches[1])

	return config.RpcUser, config.RpcPassword, nil

}


// parseConfig reads our the config settings for the
// navcoin go server and puts them into the config struct
func parseConfig(config UserConfig) UserConfig {

	config.NavConfPath = viper.GetString("navconf")
	config.RunningNavVersion = viper.GetString("runningNavVersion")

	return config

}