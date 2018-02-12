package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/NAVCoin/navpi-go/app/conf"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/NAVCoin/navpi-go/app/manager/managerapi"
	"github.com/NAVCoin/navpi-go/app/daemon/daemonapi"
	"github.com/NAVCoin/navpi-go/app/daemon"
	"github.com/NAVCoin/navpi-go/app/boxsetup/setupapi"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func world(w http.ResponseWriter, r *http.Request) {
	server.Shutdown(nil)
	fmt.Fprintf(w, "world")
}

var server *http.Server

func main() {

	//log.Println(fmt.Sprintf("Server running in %s:%s", runtime.GOOS, runtime.GOARCH))
	//log.Println(fmt.Sprintf("App pid : %d.", os.Getpid()))
	//
	serverConfig, err := conf.LoadServerConfig()
	if err != nil {
		log.Fatal("Failed to load the server config: " + err.Error())
	}

	conf.LoadUserConfig()

	router := mux.NewRouter()

	// check to see if we have a defined running config
	// If not we are only going to boot the setup apis, otherwise we will start the app
	if conf.UserConf.RunningNavVersion == "" {

		log.Println("No user config - adding setup api")
		setupapi.InitSetupHandlers(router, "api")

	} else {

		log.Println("User config found - booting all apis")

		// we have a user config so start the app in running mode
		daemon.StartManager()

		managerapi.InitManagerhandlers(router,"api")
		daemonapi.InitChainHandlers(router, "api")

	}


	// Start the server
	port := fmt.Sprintf(":%d", serverConfig.ManagerAiPort)
	srv := &http.Server{
		Addr: port,
		Handler: handlers.CORS()(router)}

	srv.ListenAndServe()

}