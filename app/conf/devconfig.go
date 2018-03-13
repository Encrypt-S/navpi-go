package conf

import "github.com/spf13/viper"

// LoadDevConfig sets up viper, loads, parses dev config
func LoadDevConfig() error {

	viper.SetConfigName("dev-config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./app")
	viper.AddConfigPath("../")

	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	parseDevConfig()

	return nil

}

// parseDevConfig parses, overrides the NavCoin conf settings
func parseDevConfig() {

	if viper.GetString("navConfig.rpcUser") != "" {
		NavConf.RPCUser = viper.GetString("navConfig.rpcUser")
	}
	if viper.GetString("navConfig.rpcPassword") != "" {
		NavConf.RPCPassword = viper.GetString("navConfig.rpcPassword")
	}

}
