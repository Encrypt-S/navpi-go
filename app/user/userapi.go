package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Encrypt-S/navpi-go/app/api"
	"github.com/Encrypt-S/navpi-go/app/conf"
	"github.com/Encrypt-S/navpi-go/app/middleware"
	"github.com/Encrypt-S/navpi-go/app/utils"
	"github.com/gorilla/mux"
)

// UIProtection defines a structure to store username and password
type LoginDetail struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// InitSetupHandlers sets the api
func InitSetupHandlers(r *mux.Router, prefix string) {

	// setup namespace
	namespace := "user"

	// login route - takes the username, password and retruns a jwt
	loginPath := api.RouteBuilder(prefix, namespace, "v1", "login")
	r.Handle(loginPath, middleware.Adapt(loginHandler(), middleware.CorsHandler()))

}

// protectUIHandler takes the api response and checks username and password
func loginHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var loginDetail LoginDetail
		apiResp := api.Response{}

		// get the json from the post data
		err := json.NewDecoder(r.Body).Decode(&loginDetail)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			returnErr := api.AppRespErrors.ServerError
			returnErr.ErrorMessage = fmt.Sprintf("Server error: %v", err)
			apiResp.Errors = append(apiResp.Errors, returnErr)
			apiResp.Send(w)

			return
		}

		// load the hash and check the details
		hashedDetails := conf.AppConf.UIPassword
		isValid := api.CheckHashDetails(loginDetail.Username, loginDetail.Password, hashedDetails)

		if !isValid {
			w.WriteHeader(http.StatusBadRequest)

			returnErr := api.AppRespErrors.LoginError
			apiResp.Errors = append(apiResp.Errors, returnErr)
			apiResp.Send(w)

			return
		}

		apiResp.Data = utils.GenerateJWT(time.Hour*24, []byte(conf.ServerConf.JWTSecret))
		apiResp.Send(w)

	})
}
