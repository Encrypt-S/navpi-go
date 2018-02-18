package daemonapi

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NAVCoin/navpi-go/app/conf"
	"github.com/NAVCoin/navpi-go/app/daemon/daemonrpc"
	"github.com/gorilla/mux"
)

// Setup all the handlers for the blockchain rpc interface
func InitWalletHandlers(r *mux.Router, prefix string) {

	r.HandleFunc(fmt.Sprintf("/%s/wallet/v1/getstakereport", prefix), geStakeReport).Methods("GET")

}

func geStakeReport(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "NAVCoin pi server") // send data to client side

	n := daemonrpc.RpcRequestData{}
	n.Method = "getstakereport"

	resp, err := daemonrpc.RequestDaemon(n, conf.NavConf)

	if err != nil { // Handle errors requesting the daemon
		daemonrpc.RpcFailed(err, w, r)
		return
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Write(bodyText)
}
