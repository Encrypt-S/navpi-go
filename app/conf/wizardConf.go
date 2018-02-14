package conf

import (
	"github.com/spf13/viper"
	"time"
)


// WizardConfigStruct contains the app's user configuration
type WizardConfigStruct struct {
	NavConfPath       string `json:"navConfPath"`
	RunningNavVersion string `json:"runningNavVersion"`
}
var WizardConf WizardConfigStruct

 func StartConfigManager() {

 	ticker := time.NewTicker(time.Millisecond * 500)
 	go func() {
 		for range ticker.C {
			LoadWizardConfig()
 		}
 	}()
 }

// SetupViper loads the userconfig from a file
func SetupViper() {

	viper.SetConfigName("wizard-config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./app")
	viper.AddConfigPath(".")
	// sviper.WatchConfig()

	//viper.OnConfigChange(func(e fsnotify.Event) {
	//ØOs
	//	fmt.Println("Config file changed:", e.Name)
	//
	//	//LoggingConfig()
	//	//WizardConf.NavCon
	//	//userconfig := parseUserConfig(WizardConfigStruct{})
	//
	//})

}

func LoadWizardConfig() error {

	err := viper.ReadInConfig()

	if err != nil {
		return err

	}
	// load user config
	userconfig := parseUserConfig(WizardConfigStruct{})

	if err != nil {
		return err
	}

	// store the user config
	WizardConf = userconfig

	return nil
}



// parseUserConfig puts settings from server-config.json into WizardConfigStruct struct
func parseUserConfig(userconfig WizardConfigStruct) WizardConfigStruct {
	userconfig.NavConfPath = viper.GetString("navconf")
	userconfig.RunningNavVersion = viper.GetString("runningNavVersion")

	return userconfig

}
