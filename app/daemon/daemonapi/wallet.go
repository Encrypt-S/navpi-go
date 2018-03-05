package daemonapi

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NAVCoin/navpi-go/app/conf"
	"github.com/NAVCoin/navpi-go/app/daemon/daemonrpc"
	"github.com/gorilla/mux"
)

// InitWalletHandlers sets up handlers for the blockchain rpc interface
func InitWalletHandlers(r *mux.Router, prefix string) {

	r.HandleFunc(fmt.Sprintf("/%s/wallet/v1/getstakereport", prefix), getStakeReport).Methods("GET")

}

// getStakeReport takes writer, request - writes out stake report
func getStakeReport(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintf(w, "NAVCoin pi server") // send data to client side

	n := daemonrpc.RpcRequestData{}
	n.Method = "getstakereport"

	resp, err := daemonrpc.RequestDaemon(n, conf.NavConf)

	// Handle errors requesting the daemon
	if err != nil {
		daemonrpc.RpcFailed(err, w, r)
		return
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Write(bodyText)
}
