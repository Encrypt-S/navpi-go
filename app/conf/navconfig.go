package conf
//
//import (
//	"errors"
//	"io/ioutil"
//	"os"
//	"regexp"
//)
//
type NavConfig struct {
	RpcUser     string `json:"rpcUser"`
	RpcPassword string `json:"rpcPassword"`
}


func CreateRPCDetails () {
	//TODO: create an random user and password on start
	NavConf.RpcPassword = "hi"
	NavConf.RpcUser = "userYo"

}


//
////func LoadRPCDetails(appconfig AppConfig) error {
////
//	var navconf = AppConf.NavConf
//
//	serverConfigFile, err := os.Open(navconf)
//	if err != nil {
//		return err
//	}
//
//	defer serverConfigFile.Close()
//	content, err := ioutil.ReadAll(serverConfigFile)
//	if err != nil {
//		return err
//	}
//
//	rpcUserRegexp, err := regexp.Compile(`(?m)^\s*rpcuser=([^\s]+)`)
//	if err != nil {
//		return err
//	}
//
//	userSubmatches := rpcUserRegexp.FindSubmatch(content)
//	if userSubmatches == nil {
//		return errors.New("No RPC User set in the config")
//	}
//
//	rpcPassRegexp, err := regexp.Compile(`(?m)^\s*rpcpassword=([^\s]+)`)
//	if err != nil {
//		return err
//	}
//
//	passSubmatches := rpcPassRegexp.FindSubmatch(content)
//	if passSubmatches == nil {
//		return errors.New("No RPC Password set")
//	}
//
//	NavConf.RpcUser = string(userSubmatches[1])
//	NavConf.RpcPassword = string(passSubmatches[1])
//
//	return nil
//
//}
