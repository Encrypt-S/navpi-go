package main

import (
	"fmt"
	"net/http"
	"io"
	"log"
	"runtime"
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

	log.Println(fmt.Sprintf("Server running in %s:%s", runtime.GOOS, runtime.GOARCH))
	log.Println(fmt.Sprintf("App pid : %d.", os.Getpid()))

	//serverMuxA := http.NewServeMux()
	//serverMuxA.HandleFunc("/hello", hello)

	serverMuxB := http.NewServeMux()
	serverMuxB.HandleFunc("/world", world)

	//server = start()


	http.ListenAndServe("localhost:8082", serverMuxB)
}

func start () *http.Server {
	srv := &http.Server{Addr: ":8081"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})

	go func() {
		srv.ListenAndServe()
		//http.ListenAndServe("localhost:8081", serverMuxA)
	}()

	return srv
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
//func populateRPCDetails(userConfig *conf.Config)  {
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
//func startAPIServer (port string, config *conf.Config) {
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
