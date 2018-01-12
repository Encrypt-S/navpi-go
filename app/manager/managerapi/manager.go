package managerapi

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"fmt"
	"io"
)





// Setup all the handlers for the blockchain rpc interface
func InitManagerhandlers(r *mux.Router, prefix string)  {


	r.HandleFunc(fmt.Sprintf("/%s/manager/v1/daemon/restart", prefix), startDaemon).Methods("GET")

}



func startDaemon(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "NAVCoin pi server") // send data to client side

	log.Println("startDaemon")



	//n := deamonrpc.RpcRequestData{}
	//n.Method = "getblockcount"
	//
	//resp, err := deamonrpc.RequestDaemon(n, config)
	//
	//if err != nil { // Handle errors requesting the daemon
	//	deamonrpc.RpcFailed(err, w, r)
	//	return
	//}
	//
	//bodyText, err := ioutil.ReadAll(resp.Body)
	//w.WriteHeader(resp.StatusCode)
	//w.Write(bodyText)
	io.WriteString(w, "Start")
}


