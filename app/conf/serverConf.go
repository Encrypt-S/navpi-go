package conf

import "github.com/spf13/viper"

// UserConfig the application's configuration
type ServerConfig struct {
	ManagerAiPort    int64
	DaemonApiPort    int64
	SetupApiPort     int64
	LatestReleaseAPI string
	ReleaseAPI string
	DaemonHeartbeat int64
}




// LoadUserConfig loads the config from a file
func LoadServerConfig() (ServerConfig, error)  {

	viper.SetConfigName("server-config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./app")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		return ServerConfig{},err
	}

	// load the go server config
	serverConfig := parseServerConfig(ServerConfig{})


	ServerConf = serverConfig

	return serverConfig, nil
}


func parseServerConfig(config ServerConfig) ServerConfig  {

	config.ManagerAiPort = viper.GetInt64("managerApiPort")
	config.LatestReleaseAPI = viper.GetString("latestReleaseAPI")
	config.ReleaseAPI = viper.GetString("releaseAPI")
	config.DaemonHeartbeat = viper.GetInt64("daemonHeartbeat")

	return config

}


