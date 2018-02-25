package setupapi

import (
	"fmt"
	"github.com/NAVCoin/navpi-go/app/middleware"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"github.com/NAVCoin/navpi-go/app/api"
	"encoding/json"
	"github.com/NAVCoin/navpi-go/app/conf"
)



func InitSetupHandlers(r *mux.Router, prefix string) {

	var nameSpace string = "setup"

	var path_ip_detect string = fmt.Sprintf("/%s/%s/v1/setrange", prefix, nameSpace)

	r.Handle(path_ip_detect, middleware.Adapt(rangeSetHandler(), middleware.Notify()))

}

func rangeSetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		host, _, err := net.SplitHostPort(r.RemoteAddr)


		apiResp := api.APIResponse{}

		// If there is no host found we need to
		if err != nil || host == "" {

			w.WriteHeader(http.StatusInternalServerError)
			apiResp.Error = api.ApiRespErrors.SetupAPINoHost

			jsonValue, _ := json.Marshal(apiResp)
			w.Write(jsonValue)
			return
		}

		if host != "::1" {

			conf.AppConf.DetectedIp = host
			conf.SaveAppConfig()

			apiResp.Success = true
			apiResp.Data = host

			jsonValue, _ := json.Marshal(apiResp)
			w.Write(jsonValue)

			return

		} else {

			w.WriteHeader(http.StatusBadRequest)
			apiResp.Error = api.ApiRespErrors.SetupAPIUsingLocalHost

			jsonValue, _ := json.Marshal(apiResp)
			w.Write(jsonValue)
			return

		}


	})
}
