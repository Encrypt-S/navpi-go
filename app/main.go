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
	conf.StartConfigManager()

	//
	//
	//
	//// Get the user config
	//// -----------------------
	////userConfig, err := conf.LoadUserConfig()
	//if err != nil {
	//	log.Fatal("Failed to load user config: " + err.Error())
	//	//startSetupApiSercer(fmt.Sprintf(":%d", serverConfig.SetupApiPort))
	//} else {
	//	// if there is no error the populate the user config
	//}
	//

	//serverMuxA := http.NewServeMux()
	//serverMuxA.HandleFunc("/hello", hello)

	// start the daemon server
	//daemonsvr.Start(serverConfig)
	//setupsrv.Start(serverConfig)


	// if we have a user config then we will start the system
	// otherwise the UI will start it later
	//if( daemonapi.UserConfig != nil) {

		//daemon.DownloadAndStart(serverConfig, daemonapi.UserConfig)
	//}

	//setupsrv.Start(serverConfig)


	daemon.StartManager()

	// start the manager server
	router := mux.NewRouter()
	managerapi.InitManagerhandlers(router,"api")
	daemonapi.InitChainHandlers(router, "api")

	port := fmt.Sprintf(":%d", serverConfig.ManagerAiPort)
	srv := &http.Server{
		Addr: port,
		Handler: handlers.CORS()(router)}

	srv.ListenAndServe()

}





//package main
//
//import (
//	"log"
//	"fmt"
//	"os"
//	"runtime"
//	"github.com/NAVCoin/navpi-go/app/conf"
//	"github.com/gorilla/mux"
//	"github.com/NAVCoin/navpi-go/app/api/blockchainapi"
//	"github.com/NAVCoin/navpi-go/app/api/addressindexapi"
//	"github.com/NAVCoin/navpi-go/app/api/walletapi"
//	"net/http"
//	"github.com/gorilla/handlers"
//	"io"
//)
//
//
//func main() {
//
//	log.Println(fmt.Sprintf("Server running in %s:%s", runtime.GOOS, runtime.GOARCH))
//	log.Println(fmt.Sprintf("App pid : %d.", os.Getpid()))
//
//
//	// Load the server config
//	//-----------------------
//	serverConfig, err := conf.LoadServerConfig()
//	if err != nil {
//		log.Fatal("Failed to load the server config: " + err.Error())
//	}
//
//	// Get the user config
//	//-----------------------
//	userConfig, err := conf.LoadUserConfig()
//
//
//
//	if err != nil {
//		//log.Fatal("Failed to load config: " + err.Error())
//		//startSetupApiSercer(fmt.Sprintf(":%d", serverConfig.SetupApiPort))
//
//	}
//
//
//
//	srv := startHttpServer()
//
//
//
//
//	// Get the RPC details
//	//-----------------------------------
//
//
//	// check the daemon and path
//	//daemonPath, err := daemon.CheckDaemon(serverConfig)
//	//if err != nil {
//	//	log.Fatal("Failed on checking the daemon error " + err.Error())
//	//}
//
//	// load both the user settings and nav coin configs
//	//config, err := conf.LoadUserConfig()
//	//if err != nil {
//	//	log.Fatal("Failed to load config: " + err.Error())
//	//}
//	//
//	//
//	//log.Println( fmt.Sprintf("Straring server on port :%d", serverConfig.Port))//read the server port
//	//port := fmt.Sprintf(":%d", serverConfig.Port)
//	//
//	//startAPIServer(port, config)
//
//}
//
//
//
//// loads the RPC details from the path given in the config
//func populateRPCDetails(userConfig *conf.UserConfig)  {
//	// we have the user config soe
//	rpcUser, rpcPassword, err := conf.LoadRPCDetails(userConfig)
//	if err != nil {
//		log.Fatal("Failed to get rpc details: " + err.Error())
//	}
//	userConfig.RpcUser = rpcUser
//	userConfig.RpcPassword = rpcPassword
//}
//
//
////func startSetupApiSercer(port string) {
////	startHttpServer(port)
////}
//
//func startAPIServer (port string, config *conf.UserConfig) {
//
//	router := mux.NewRouter()
//
//
//
//	//add all the apis
//	blockchainapi.InitHandlers(router, config, "api")
//	addressindexapi.InitHandlers(router, config, "api")
//	walletapi.InitHandlers(router, config, "api")
//
//
//
//	log.Fatal(http.ListenAndServe(port, handlers.CORS()(router)))
//}
//
//
//func startHttpServer() *http.Server {
//	srv := &http.Server{Addr: ":8080"}
//
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		io.WriteString(w, "hello world\n")
//	})
//
//	go func() {
//		if err := srv.ListenAndServe(); err != nil {
//			// cannot panic, because this probably is an intentional close
//			log.Printf("Httpserver: ListenAndServe() error: %s", err)
//		}
//	}()
//
//	// returning reference so caller can call Shutdown()
//	return srv
//}
