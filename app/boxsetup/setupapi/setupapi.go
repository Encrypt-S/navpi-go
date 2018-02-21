package setupapi

import (
	"fmt"
	"github.com/NAVCoin/navpi-go/app/conf"
	"github.com/NAVCoin/navpi-go/app/middleware"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
)

func InitSetupHandlers(r *mux.Router, prefix string) {

	var nameSpace string = "setup"

	// test @ http://192.168.1.207:9002/api/setup/v1/detectip
	var path_ip_detect string = fmt.Sprintf("/%s/%s/v1/detectip", prefix, nameSpace)

	r.Handle(path_ip_detect, middleware.Adapt(detectIpV1Handler(), middleware.Notify()))

}

func detectIpV1Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		host, port, err := net.SplitHostPort(r.RemoteAddr)

		// host :: 192.168.1.207
		log.Println("host :: ", host)

		// port :: 59660
		 log.Println("port :: ", port)

		if err != nil || host == "" {
			log.Println("err :: ", err)
		}

		if host == "::1" {

			log.Println("localhost")

		} else {

			log.Println("not localhost")

		}

		conf.AppConf.DetectedIp = host
		conf.SaveAppConfig()

		fmt.Fprintf(w, "Hi there, I ran the middleware, I love %s!", r.URL.Path[1:])

	})
}
