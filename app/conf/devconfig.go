package conf

import "github.com/spf13/viper"

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

func parseDevConfig() {

	// Override the NavCoin conf settings
	if viper.GetString("navConfig.rpcUser") != "" {
		NavConf.RpcUser = viper.GetString("navConfig.rpcUser")
	}

	if viper.GetString("navConfig.rpcPassword") != "" {
		NavConf.RpcUser = viper.GetString("navConfig.rpcPassword")
	}

}
