package addressindexapi

import (
	"github.com/NAVCoin/navpi-go/app/conf"
	"github.com/gorilla/mux"
)


type Resp struct  {
	Code int
	Data string
	Message string
}

var config *conf.Config


// Setup all the handlers for the blockchain rpc interface
func InitHandlers(r *mux.Router, conf *conf.Config, prefix string )  {

	config = conf
	println(prefix)


	//r.HandleFunc("/addressindex/v1/getaddressbalance", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/addressindex/v1/getaddressdeltas", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/addressindex/v1/getaddressmempool", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/addressindex/v1/getaddresstxids", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/addressindex/v1/getaddressutxos", api.NotImplemented).Methods("GET")

}



