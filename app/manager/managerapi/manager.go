package managerapi

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

// Setup all the handlers for the blockchain rpc interface
func InitManagerhandlers(r *mux.Router, prefix string) {

	r.HandleFunc(fmt.Sprintf("/%s/manager/v1/daemon/restart", prefix), startDaemon).Methods("GET")

}

func startDaemon(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "NAVCoin pi server") // send data to client side

	log.Println("resart daemon requested")

	//n := daemonrpc.RpcRequestData{}
	//n.Method = "getblockcount"
	//
	//resp, err := daemonrpc.RequestDaemon(n, config)
	//
	//if err != nil { // Handle errors requesting the daemon
	//	daemonrpc.RpcFailed(err, w, r)
	//	return
	//}
	//
	//bodyText, err := ioutil.ReadAll(resp.Body)
	//w.WriteHeader(resp.StatusCode)
	//w.Write(bodyText)
	io.WriteString(w, "Start")
}
