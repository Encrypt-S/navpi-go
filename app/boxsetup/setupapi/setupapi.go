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
	"github.com/muesli/crunchy"
)

type UIProtection struct {

	Username string `json:"username"`
	Password string `json:"password"`

}


// InitSetupHandlers sets the api
func InitSetupHandlers(r *mux.Router, prefix string) {

	var nameSpace string = "setup"

	r.Handle(fmt.Sprintf("/%s/%s/v1/setrange", prefix, nameSpace), middleware.Adapt(rangeSetHandler(), middleware.Notify()))

	// Protect UI with username and password
	r.Handle(fmt.Sprintf("/%s/%s/v1/protectui", prefix, nameSpace), middleware.Adapt(protectUIHandler())).Methods("POST")

}




// rangeSetHandler takes the users ip address and saves it to the config as a range
func protectUIHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var uiProtection UIProtection
		//params := mux.Vars(r)
		err := json.NewDecoder(r.Body).Decode(&uiProtection)

		apiResp := api.Response{}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			apiResp.Error = api.AppRespErrors.ServerError
		}

		if uiProtection.Username == "" || uiProtection.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			apiResp.Error = api.AppRespErrors.SetupAPIProtectUI

		}

		validator := crunchy.NewValidator()
		err = validator.Check(uiProtection.Password)

		if err != nil {
			fmt.Printf("The password '%s' is considered unsafe: %v\n", uiProtection.Password, err)
		}

		// has the details for later
		hashedDetails, err := api.HashDetails(uiProtection.Username, uiProtection.Password)

		// if there was an error hasing the details then error
		if err != nil {

			w.WriteHeader(http.StatusInternalServerError)
			apiResp.Error = api.AppRespErrors.ServerError

		} else {

			//update the uihash
			conf.AppConf.UIPassword = hashedDetails
			conf.SaveAppConfig()

			apiResp.Success = true

		}

		jsonValue, _ := json.Marshal(apiResp)
		w.Write(jsonValue)

	})
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


