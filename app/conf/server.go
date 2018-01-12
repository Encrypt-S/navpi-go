package conf

import "github.com/spf13/viper"

// Config the application's configuration
type ServerConfig struct {
	DaemonApiPort      	int64
	SetupApiPort      	int64
	ReleaseAPI 	string
}




// LoadUserConfig loads the config from a file
func LoadServerConfig() (*ServerConfig, error)  {

	viper.SetConfigName("server-config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./app")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		return nil,err
	}

	// load the go server config
	serverConfig := new(ServerConfig)
	parseServerConfig(serverConfig)

	return serverConfig, nil
}


func parseServerConfig(config *ServerConfig)  {

	config.DaemonApiPort = viper.GetInt64("daemonApiPort")
	config.SetupApiPort = viper.GetInt64("setupApiPort")
	config.ReleaseAPI = viper.GetString("releaseAPI")

}

