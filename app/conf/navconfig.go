package conf

import "github.com/NAVCoin/navpi-go/app/utils"

// NavConfig defines a structure to store rpc data
type NavConfig struct {
	RPCUser     string `json:"rpcUser"`
	RPCPassword string `json:"rpcPassword"`
}

// CreateRPCDetails generate rpc details for this run
func CreateRPCDetails() {

	NavConf.RPCUser, _ = utils.GenerateRandomString(32)
	NavConf.RPCPassword, _ = utils.GenerateRandomString(32)

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
//	NavConf.RPCUser = string(userSubmatches[1])
//	NavConf.RPCPassword = string(passSubmatches[1])
//
//	return nil
//
//}
