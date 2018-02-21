package setupapi

import (
	"fmt"
	"github.com/NAVCoin/navpi-go/app/middleware"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"github.com/NAVCoin/navpi-go/app/conf"
)

// Setup all the handlers for the blockchain rpc interface
func InitSetupHandlers(r *mux.Router, prefix string) {

	var nameSpace string = "setup"

	// hello route path
	//var path_hello string = fmt.Sprintf("/%s/%s/v1/hello", prefix, nameSpace)

	// notify route path
	//var path_notify string = fmt.Sprintf("/%s/%s/v1/notify", prefix, nameSpace)

	// whitelist route path
	//var path_whitelist string = fmt.Sprintf("/%s/%s/v1/whitelist", prefix, nameSpace)

	var path_ip_detect string = fmt.Sprintf("/%s/%s/v1/detectip", prefix, nameSpace)

	// standard hello world without middleware
	//r.HandleFunc(path_hello, hello).Methods("GET")

	// whitelist route using adapter middleware
	//r.Handle(path_whitelist, middleware.Adapt(whitelistV1Handler(), middleware.Notify()))

	// ipDetection route using adapter middleware
	r.Handle(path_ip_detect, middleware.Adapt(ipDetectV1Handler(), middleware.Notify()))

	// notify route using adapter middleware
	//r.Handle(path_notify, middleware.Adapt(notifyV1Handler(), middleware.Notify()))

}

// whitelist route handler
func ipDetectV1Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("ipDetectV1Handler r=", r)

		host, port, err := net.SplitHostPort(r.RemoteAddr)

		log.Println("host=", host)
		log.Println("port=", port)
		log.Println("err=", err)

		if err != nil || host == "" {
			//w.WriteHeader(http.StatusInternalServerError)
		}

		if host == "::1" {
			log.Println("we are on localhost!")
		}

		// if we are not in localhost parse the IP
		if host != "::1" {

			log.Println("we are not on localhost")

			//mockHost := net.ParseIP("51.1.1.10")

			//requestIP := net.ParseIP(host)

			//log.Println(mockHost)

			//conf.AppConf.DetectedIp = mockHost

			// now save the ip to AppConfig
			//conf.SaveAppConfig()

		}

		fmt.Fprintf(w, "Hi there, I ran the middleware, I love %s!", r.URL.Path[1:])
	})
}

// notify handler
//func notifyV1Handler() http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Println("notifyV1Handler")
//		fmt.Fprintf(w, "Hi there, I ran the middleware, I love %s!", r.URL.Path[1:])
//	})
//}


// hello world handler
//func hello(w http.ResponseWriter, r *http.Request) {

	//log.Println("hello")

	//n := daemonrpc.RpcRequestData{}
	//n.Method = "getblockcount"
	//
	//resp, err := daemonrpc.RequestDaemon(n, config)
	//
	//if err != nil { // Handle errors requesting the daemon
	//	daemonrpc.RpcFailed(err, w, r)
	//	return
	//}
	//
	//bodyText, err := ioutil.ReadAll(resp.Body)
	//w.WriteHeader(resp.StatusCode)
	//w.Write(bodyText)

	//io.WriteString(w, "Hello ")

//}
