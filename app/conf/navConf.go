package conf

import (
	"regexp"
	"os"
	"io/ioutil"
	"errors"
)

// NavConfig contains navcoin.conf's configuration
type NavConfig struct {
	RpcUser     string
	RpcPassword string
}

var NavConf NavConfig

// LoadRPCDetails tries to read the config file for the RPC server
// and extract the RPC user and password from it.
func LoadRPCDetails() error {

	NavConf := NavConfig{}

	userconfigfile := WizardConf.NavConfPath

	// Read the RPC server userconfig
	serverConfigFile, err := os.Open(userconfigfile)
	if err != nil {
		return err
	}

	defer serverConfigFile.Close()
	content, err := ioutil.ReadAll(serverConfigFile)
	if err != nil {
		return err
	}

	// Extract the rpcuser
	rpcUserRegexp, err := regexp.Compile(`(?m)^\s*rpcuser=([^\s]+)`)
	if err != nil {
		return err
	}
	userSubmatches := rpcUserRegexp.FindSubmatch(content)
	if userSubmatches == nil {
		// No user found, nothing to do
		return errors.New("No RPC User set in the userconfig")
	}

	// Extract the rpcpass
	rpcPassRegexp, err := regexp.Compile(`(?m)^\s*rpcpassword=([^\s]+)`)
	if err != nil {
		return err
	}

	passSubmatches := rpcPassRegexp.FindSubmatch(content)


	if passSubmatches == nil {
		// No password found we will die
		return errors.New("No RPC Password set")
	}

	// save the RPCUser and RPCpassword into the navconfig
	//navconfig.RpcUser = string(userSubmatches[1])
	//navconfig.RpcPassword = string(passSubmatches[1])
	NavConf.RpcUser = string(userSubmatches[1])
	NavConf.RpcPassword = string(passSubmatches[1])

	return nil

}
