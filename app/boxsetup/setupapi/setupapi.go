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
	"strings"
)

// InitSetupHandlers sets the api
func InitSetupHandlers(r *mux.Router, prefix string) {

	var nameSpace string = "setup"

	var path_ip_detect string = fmt.Sprintf("/%s/%s/v1/setrange", prefix, nameSpace)

	r.Handle(path_ip_detect, middleware.Adapt(rangeSetHandler(), middleware.Notify()))

}

// rangeSetHandler takes the users ip address and saves it to the config as a range
func rangeSetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		host, _, err := net.SplitHostPort(r.RemoteAddr)
		apiResp := api.Response{}

		// If there is no host found we need to error out
		if err != nil || host == "" {

			w.WriteHeader(http.StatusInternalServerError)
			apiResp.Error = api.AppRespErrors.SetupAPINoHost

			jsonValue, _ := json.Marshal(apiResp)
			w.Write(jsonValue)
			return
		}


		// Note "::1"  is the ipV6 version of localhost
		// Check to see we are not using "localhost" - we need an ip
		if host == "::1" {

			w.WriteHeader(http.StatusBadRequest)
			apiResp.Error = api.AppRespErrors.SetupAPIUsingLocalHost

			jsonValue, _ := json.Marshal(apiResp)
			w.Write(jsonValue)
			return

		}

		// we made it here so we are good - so set the config and save to the file

		// separate and make the range wildcard
		strSplit := strings.Split(host, ".")
		strSplit[len(strSplit) -1 ] = "*"
		strSplit[len(strSplit) -2 ] = "*"
		host = strings.Join(strSplit, ".")


		conf.AppConf.AllowedIps = append(conf.AppConf.AllowedIps, host)
		conf.SaveAppConfig()

		//Set the rep data
		apiResp.Success = true
		apiResp.Data = host

		jsonValue, _ := json.Marshal(apiResp)
		w.Write(jsonValue)

	})
}
