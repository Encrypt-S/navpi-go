package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NAVCoin/navpi-go/app/middleware"
	"github.com/gorilla/mux"
)

// Response is the generic resp that will be used for the api
type Response struct {
	Data   interface{} `json:"data,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
	Errors []errorCode `json:"errors,omitempty"`
}

// Send marshal the response and write value
func (i *Response) Send(w http.ResponseWriter) {
	jsonValue, _ := json.Marshal(i)
	w.Write(jsonValue)
}

type errorCode struct {
	Code         string `json:"code,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type appErrorsStruct struct {
	ServerError             	errorCode
	InvalidPasswordStrength 	errorCode

	SetupAPIUsingLocalHost 	errorCode
	SetupAPINoHost         	errorCode
	SetupAPIProtectUI      	errorCode

	LoginError				errorCode
}

// AppRespErrors variable
var AppRespErrors appErrorsStruct

// BuildAppErrors builds all the error messages that the app
func BuildAppErrors() {

	AppRespErrors = appErrorsStruct{}

	// Generic errors
	AppRespErrors.ServerError = errorCode{"SERVER_ERROR", "There was an unexpected error - please try again"}
	AppRespErrors.InvalidPasswordStrength = errorCode{"INVALID_PASSWORD_STRENGTH", ""}

	// Setup API Errors
	AppRespErrors.SetupAPIUsingLocalHost = errorCode{"SETUP_HOST_NOT_FOUND", "The host was not found"}
	AppRespErrors.SetupAPINoHost = errorCode{"SETUP_USING_LOCAL_HOST", "You are using localhost, please use 127.0.01 or your network ip address"}
	AppRespErrors.SetupAPIProtectUI = errorCode{"SETUP_MISSING_USERNAME_PASSWORD", "You are missing the username and/or password"}


	AppRespErrors.LoginError = errorCode{"LOGIN_ERROR", "Your username and/or password is wrong"}
}

// InitMetaHandlers starts the meta api handlers
func InitMetaHandlers(r *mux.Router, prefix string) {

	nameSpace := "meta"

	r.Handle(fmt.Sprintf("/%s/%s/v1/errorcode", prefix, nameSpace), middleware.Adapt(metaErrorDisplayHandler()))

}

// metaErrorDisplayHandler displays all the application errors to frontend
func metaErrorDisplayHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		appResp := Response{}
		appResp.Data = AppRespErrors
		appResp.Send(w)

	})
}
