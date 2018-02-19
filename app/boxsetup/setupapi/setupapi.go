package setupapi

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"io"
	"log"
	"github.com/NAVCoin/navpi-go/app/boxsetup/netconfig"
)

// Setup all the handlers for the blockchain rpc interface
func InitSetupHandlers(r *mux.Router, prefix string)  {

	var nameSpace string = "setup"

	// hello world route
	var path_hello string = fmt.Sprintf("/%s/%s/v1/hello", prefix, nameSpace)

	// netconfig route
	var path_netconfig string = fmt.Sprintf("/%s/%s/v1/netconfig", prefix, nameSpace)

	// handle hello world
	r.HandleFunc(path_hello, hello).Methods("GET")

	// handle netconfig
	r.HandleFunc(path_netconfig, netconfig.HttpScan).Methods("GET")

}

func hello(w http.ResponseWriter, r *http.Request) {

	log.Println("hello")

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

	io.WriteString(w, "Hello ")

}




