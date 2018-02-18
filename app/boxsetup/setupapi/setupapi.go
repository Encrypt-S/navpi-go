package setupapi



import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"io"
	"log"
)

// Setup all the handlers for the blockchain rpc interface
func InitSetupHandlers(r *mux.Router, prefix string)  {

	var nameSpace string = "setup"

	var path string = fmt.Sprintf("/%s/%s/v1/hello", prefix, nameSpace)

	r.HandleFunc(path, hello).Methods("GET")

}


func hello(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "NAVCoin pi server") // send data to client side

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




