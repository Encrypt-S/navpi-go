package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NAVCoin/navpi-go/app/boxsetup/setupapi"
	"github.com/NAVCoin/navpi-go/app/conf"
	"github.com/NAVCoin/navpi-go/app/daemon"
	"github.com/NAVCoin/navpi-go/app/daemon/daemonapi"
	"github.com/NAVCoin/navpi-go/app/manager/managerapi"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"os"
	"runtime"
)

var server *http.Server

func main() {

	// log out server runtime OS and Architecture
	log.Println(fmt.Sprintf("Server running in %s:%s", runtime.GOOS, runtime.GOARCH))
	log.Println(fmt.Sprintf("App pid : %d.", os.Getpid()))

	// load the server config - this is required otherwise we die right here
	serverConfig, err := conf.LoadServerConfig()
	if err != nil {
		log.Fatal("Failed to load the server config: " + err.Error())
	}

	//conf.InitAppConfig()
	conf.LoadAppConfig()
	conf.StartConfigManager()

	router := mux.NewRouter()

	// check to see if we have a defined running config
	// If not we are only going to boot the setup apis, otherwise we will start the app
	if conf.AppConf.RunningNavVersion == "" {

		log.Println("No user config - adding setup api")
		setupapi.InitSetupHandlers(router, "api")

	} else {

		log.Println("User config found - booting all apis")

		err := conf.LoadRPCDetails(conf.AppConf)

		if err != nil {
			//TODO: Fix this
			log.Println("RPC Details Not found!")
			log.Println("err", err)
		}

		// we have a user config so start the app in running mode
		daemon.StartManager()

		managerapi.InitManagerhandlers(router, "api")
		daemonapi.InitChainHandlers(router, "api")

	}

	// Start the server
	port := fmt.Sprintf(":%d", serverConfig.ManagerApiPort)
	srv := &http.Server{
		Addr:    port,
		Handler: handlers.CORS()(router)}

	srv.ListenAndServe()

}
