package daemonapi

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"fmt"
	"github.com/NAVCoin/navpi-go/app/daemon/deamonrpc"
	"github.com/NAVCoin/navpi-go/app/conf"
)

// Setup all the handlers for the blockchain rpc interface
func InitWalletHandlers(r *mux.Router, prefix string)  {

	r.HandleFunc(fmt.Sprintf("/%s/wallet/v1/getstakereport", prefix), geStakeReport).Methods("GET")

}


func geStakeReport(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "NAVCoin pi server") // send data to client side

	n := deamonrpc.RpcRequestData{}
	n.Method = "getstakereport"

	resp, err := deamonrpc.RequestDaemon(n, conf.UserConf)

	if err != nil { // Handle errors requesting the daemon
		deamonrpc.RpcFailed(err, w, r)
		return
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Write(bodyText)
}




