package daemonsrv

import (
	"github.com/NAVCoin/navpi-go/app/conf"
	"io"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"github.com/gorilla/handlers"
	"github.com/NAVCoin/navpi-go/app/daemon/daemonapi"
)

var daemonServer *http.Server
var config *conf.Config


func Start (serverConfig *conf.ServerConfig) *http.Server {

	port := fmt.Sprintf(":%d", serverConfig.DaemonApiPort)


	router := mux.NewRouter()
	daemonapi.InitChainHandlers(router,"api")


	srv := &http.Server{
		Addr: port,
		Handler: handlers.CORS()(router)}

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	io.WriteString(w, "hello world\n")
	//})


	log.Fatal(http.ListenAndServe(port, handlers.CORS()(router)))

	go func() {
		srv.ListenAndServe()
		//http.ListenAndServe("localhost:8081", serverMuxA)
	}()

	// store it so we can get it later
	daemonServer = srv
	return srv
}
