package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Encrypt-S/navpi-go/app/middleware"
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

// appErrorsStruct defines errors, errorCodes
type appErrorsStruct struct {
	LoginError             errorCode
	InvalidStrength        errorCode
	ServerError            errorCode
	SetupAPIUsingLocalHost errorCode
	SetupAPINoHost         errorCode
	SetupAPIProtectUI      errorCode
}

// AppRespErrors variable
var AppRespErrors appErrorsStruct

// BuildAppErrors builds all the error messages that the app
func BuildAppErrors() {

	AppRespErrors = appErrorsStruct{}

	// Generic errors
	AppRespErrors.ServerError = errorCode{"SERVER_ERROR", "There was an unexpected error - please try again"}
	AppRespErrors.InvalidStrength = errorCode{"INVALID_STRENGTH", ""}

	// Setup API Errors
	AppRespErrors.SetupAPIUsingLocalHost = errorCode{"SETUP_HOST_NOT_FOUND", "The host was not found"}
	AppRespErrors.SetupAPINoHost = errorCode{"SETUP_USING_LOCAL_HOST", "You are using localhost, please use 127.0.01 or your network ip address"}
	AppRespErrors.SetupAPIProtectUI = errorCode{"SETUP_MISSING_USERNAME_PASSWORD", "You are missing the username and/or password"}

	// Login Errors
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

func RouteBuilder(prefix string, namespace string, version string, method string) string {
	route := fmt.Sprintf("/%s/%s/%s/%s", prefix, namespace, version, method)
	log.Println(route)
	return route
}

func OpenRouteHandler(path string, r *mux.Router,  f http.Handler, method string) {
	r.Handle(path, middleware.Adapt(f, middleware.CORSHandler())).Methods(method)
}

func ProtectedRouteHandler(path string, r *mux.Router,  f http.Handler,  method string ) {
	r.Handle(path, middleware.Adapt(f,
		middleware.CORSHandler(),
		middleware.JwtHandler())).
		Methods(method)
}
