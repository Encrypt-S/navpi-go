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

	var path_ip_detect string = fmt.Sprintf("/%s/%s/v1/detectip", prefix, nameSpace)

	r.Handle(path_ip_detect, middleware.Adapt(detectIpV1Handler(), middleware.Notify()))

}

func detectIpV1Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("detectIpV1Handler r=", r)

		host, port, err := net.SplitHostPort(r.RemoteAddr)

		log.Println("host=", host)
		log.Println("port=", port)
		log.Println("err=", err)

		if err != nil || host == "" {
			log.Println("err=", err)
		}

		if host == "::1" {

			log.Println("we are on localhost!")

		} else {

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
