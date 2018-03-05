package conf

import (
	"crypto/rand"
	"encoding/base64"
)

type NavConfig struct {
	RpcUser     string `json:"rpcUser"`
	RpcPassword string `json:"rpcPassword"`
}


// CreateRPCDetails generate rpc details for this run
func CreateRPCDetails () {

	NavConf.RpcUser, _ = generateRandomString(32)
	NavConf.RpcPassword, _ = generateRandomString(32)

}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
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
