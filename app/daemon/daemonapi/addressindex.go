package daemonapi

import (
	"github.com/gorilla/mux"
)

// Resp defines a structure to store response data
type Resp struct {
	Code    int
	Data    string
	Message string
}

// InitAddressHandlers sets up handlers for the blockchain rpc interface
func InitAddressHandlers(r *mux.Router, prefix string) {

	println(prefix)

	//r.HandleFunc("/addressindex/v1/getaddressbalance", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/addressindex/v1/getaddressdeltas", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/addressindex/v1/getaddressmempool", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/addressindex/v1/getaddresstxids", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/addressindex/v1/getaddressutxos", api.NotImplemented).Methods("GET")

}
