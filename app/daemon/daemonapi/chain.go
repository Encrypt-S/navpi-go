package daemonapi

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"io"
	"io/ioutil"

	"github.com/Encrypt-S/navpi-go/app/api"
	"github.com/Encrypt-S/navpi-go/app/conf"
	"github.com/Encrypt-S/navpi-go/app/daemon/daemonrpc"
)

// InitChainHandlers sets up handlers for the blockchain rpc interface
func InitChainHandlers(r *mux.Router, prefix string) {

	namespace := "chain"

	getBlocCountPath := api.RouteBuilder(prefix, namespace, "v1", "getblockcount")
	api.ProtectedRouteHandler(getBlocCountPath, r, getBlockCount(), http.MethodPost)


	//r.Handle(getBlocCountPath, middleware.Adapt(getBlockCount(),
	//	middleware.JwtHandler())).Methods("GET")

	//// not implemented
	//r.HandleFunc("/blockchain/v1/getbestblockhash", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getblock", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getblockchaininfo", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getblockhash", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getblockhashes", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getblockheader", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getchaintips", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getdifficulty", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getmempoolancestors", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getmempoolentry", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getmempoolinfo", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getrawmempool", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getspentinfo", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/gettxout", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/gettxoutproof", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/gettxoutsetinfo", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/verifychain", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/verifytxoutproof", api.NotImplemented).Methods("GET")

}

// protectUIHandler takes the api response and checks username and password
func getBlockCount() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("getBlockCount")

		n := daemonrpc.RpcRequestData{}
		n.Method = "getblockcount"

		resp, err := daemonrpc.RequestDaemon(n, conf.NavConf)

		if err != nil { // Handle errors requesting the daemon
			daemonrpc.RpcFailed(err, w, r)
			return
		}

		bodyText, err := ioutil.ReadAll(resp.Body)
		w.WriteHeader(resp.StatusCode)
		w.Write(bodyText)
		io.WriteString(w, "hello world\n")
	})
}
