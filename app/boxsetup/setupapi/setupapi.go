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
	io.WriteString(w, "Hello ")

}




