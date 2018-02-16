package conf

import "github.com/spf13/viper"

// SetupViper loads the userconfig from a file
func SetupViper() {

	viper.SetConfigName("user-config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./app")
	viper.AddConfigPath(".")
}
