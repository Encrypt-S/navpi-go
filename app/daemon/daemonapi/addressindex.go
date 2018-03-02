package daemonapi

import (
	"github.com/gorilla/mux"
)

type Resp struct {
	Code    int
	Data    string
	Message string
}

// Setup all the handlers for the blockchain rpc interface
func InitAddressHandlers(r *mux.Router, prefix string) {

	println(prefix)

	//r.HandleFunc("/addressindex/v1/getaddressbalance", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/addressindex/v1/getaddressdeltas", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/addressindex/v1/getaddressmempool", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/addressindex/v1/getaddresstxids", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/addressindex/v1/getaddressutxos", api.NotImplemented).Methods("GET")

}
