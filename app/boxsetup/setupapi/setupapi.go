package setupapi

import (
	"fmt"
	"github.com/NAVCoin/navpi-go/app/middleware"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"github.com/NAVCoin/navpi-go/app/api"
	"encoding/json"
)



func InitSetupHandlers(r *mux.Router, prefix string) {

	var nameSpace string = "setup"

	var path_ip_detect string = fmt.Sprintf("/%s/%s/v1/setrange", prefix, nameSpace)

	r.Handle(path_ip_detect, middleware.Adapt(rangeSetHandler(), middleware.Notify()))

}

func rangeSetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		host, _, err := net.SplitHostPort(r.RemoteAddr)

		w.WriteHeader(http.StatusBadRequest)

		// If there is no host found we need to
		if err != nil || host == "" {
			apiResp := api.APIResponse{}
			apiResp.Error.Message = "Cannot detect the host"

			jsonValue, _ := json.Marshal(apiResp)

			w.WriteHeader(http.StatusInternalServerError)s
			w.Write(jsonValue)

			return
		}

		if host != "::1" {


		}

		//host, port, err := net.SplitHostPort(r.RemoteAddr)
		//
		//log.Println("host :: ", host)
		//
		//log.Println("port :: ", port)
		//
		//if err != nil || host == "" {
		//	log.Println("err :: ", err)
		//}
		//
		//if host == "::1" {
		//
		//	log.Println("localhost")
		//
		//} else {
		//
		//	log.Println("not localhost")
		//
		//}
		//
		//conf.AppConf.DetectedIp = host
		//conf.SaveAppConfig()
		//
		//fmt.Fprintf(w, "Hi there, I ran the middleware, I love %s!", r.URL.Path[1:])

	})
}
