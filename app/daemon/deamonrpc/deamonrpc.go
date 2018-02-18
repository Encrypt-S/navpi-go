package deamonrpc

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/NAVCoin/navpi-go/app/conf"
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

	username := navConf.RpcUser
	password := navConf.RpcPassword

	client := &http.Client{}

	jsonValue, _ := json.Marshal(rpcReqData)

	url := "http://127.0.0.1:44444"

	//if(config.TestNet) {
	//	url = "http://127.0.0.1:44445"
	//}

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

	//Write json response back to response
	w.Write(respJson)

}

// NotImplemented: this is a generic function for daeomn apis
// that have not been implemented yet
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

	//Write json response back to response
	w.Write(respJson)

}
