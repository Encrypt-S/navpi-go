package daemonrpc

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/NAVCoin/navpi-go/app/conf"
	"fmt"
)

type RpcRequestData struct {
	Method string `json:"method"`
}

type RpcResp struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

// RequestDaemon request the data via the daemon's rpc api
// it also allows auto switches between the testnet and live depending on the config
func RequestDaemon(rpcReqData RpcRequestData, navConf conf.NavConfig) (*http.Response, error) {

	serverConf := conf.ServerConf

	username := navConf.RPCUser
	password := navConf.RPCPassword

	client := &http.Client{}

	jsonValue, _ := json.Marshal(rpcReqData)

	// set the port to live
	port := serverConf.LivePort

	// check to see if we are in test net mode
	if serverConf.UseTestnet {
		port = serverConf.TestPort
	}

	// build the url
	url := fmt.Sprintf("http://127.0.0.1:%d", port)

	//  build the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.SetBasicAuth(username, password)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	return resp, err

}

func RpcFailed(err error, w http.ResponseWriter, r *http.Request) {

	resp := RpcResp{}

	w.WriteHeader(http.StatusFailedDependency)
	resp.Code = http.StatusFailedDependency
	resp.Message = "Failed to run command: " + err.Error()
	log.Fatal("Failed to run command: " + err.Error())

	respJson, err := json.Marshal(resp)

	if err != nil {

	}

	w.Write(respJson)

}

// NotImplemented: this is a generic function for daeomn apis
func NotImplemented(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "NAVCoin pi server") // send data to client side

	w.Header().Set("Content-Type", "application/json")
	resp := RpcResp{}

	w.WriteHeader(http.StatusNotImplemented)
	resp.Code = http.StatusNotImplemented
	resp.Message = "Not implemented"

	respJson, err := json.Marshal(resp)

	if err != nil {

	}

	w.Write(respJson)

}
