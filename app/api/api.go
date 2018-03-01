package api

import (
	"github.com/gorilla/mux"
	"fmt"
	"github.com/NAVCoin/navpi-go/app/middleware"
	"encoding/json"
	"net/http"
)

// The generic resp that will be used for the api
type Response struct {
	Data    interface{}    	`json:"data,omitempty"`
	Meta 	interface{} 	`json:"meta,omitempty"`
	Errors	[]errorCode 	`json:"errors,omitempty"`
}

// Send marshal the response and writes it our
func (i *Response) Send (w http.ResponseWriter) {
	jsonValue, _ := json.Marshal(i)
	w.Write(jsonValue)
}


type errorCode struct {
	 Code string 		`json:"code,omitempty"`
	 ErrorMessage string `json:"errorMessage,omitempty"`
}


type appErrorsStruct struct {

	ServerError             	errorCode
	InvalidPasswordStrength 	errorCode

	SetupAPIUsingLocalHost 	errorCode
	SetupAPINoHost         	errorCode
	SetupAPIProtectUI 		errorCode

}


var AppRespErrors appErrorsStruct




/**
Build errors builds all the error messages that the app
will use and display to the error.
 */
func BuildAppErrors()  {

	AppRespErrors = appErrorsStruct{}

	// Generic errors
	AppRespErrors.ServerError = errorCode{"SERVER_ERROR", "There was an unexpected error - please try again"}
	AppRespErrors.InvalidPasswordStrength = errorCode{"INVALID_PASSWORD_STRENGTH", ""}

	// Setup API Errors
	AppRespErrors.SetupAPIUsingLocalHost = errorCode{"SETUP_HOST_NOT_FOUND", "The host was not found"}
	AppRespErrors.SetupAPINoHost = errorCode{"SETUP_USING_LOCAL_HOST", "You are using localhost, please use 127.0.01 or your network ip address"}
	AppRespErrors.SetupAPIProtectUI = errorCode{"SETUP_MISSING_USERNAME_PASSWORD", "You are missing the username and/or password"}

}


// Starts the meta api handlers
func InitMetaHandlers(r *mux.Router, prefix string) {

	nameSpace := "meta"

	r.Handle( fmt.Sprintf("/%s/%s/v1/errorcode", prefix, nameSpace), middleware.Adapt(metaErrorDisplayHandler()))

}



// metaErrorDisplayHandler displays all the application errors
// to the frontend
func metaErrorDisplayHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		appResp := Response{}
		appResp.Data = AppRespErrors
		appResp.Send(w)

	})
}
