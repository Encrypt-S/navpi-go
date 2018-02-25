package api

import (
	"github.com/gorilla/mux"
)

type APIResponse struct {

	Data string `json:"data"`
	Success bool `json:"success"`
	Error ErrorCode `json:"error"`

}


type ErrorCode struct {

	 code string `json:"code"`
	  ErrorMessage string `json:"errorMessage"`
}


type AppErrorsStruct struct {
	SetupAPIUsingLocalHost ErrorCode
	SetupAPINoHost ErrorCode

}


var ApiRespErrors AppErrorsStruct


func BuildAppErrors()  {

	AppErrors := AppErrorsStruct{}

	AppErrors.SetupAPIUsingLocalHost = ErrorCode{"SETUP_HOST_NOT_FOUND", "The host was not found"}
	AppErrors.SetupAPINoHost = ErrorCode{"USING_LOCAL_HOST", "You are using localhost, please use 127.0.01 or your network ip address"}

}



func InitMetaHandlers(r *mux.Router, prefix string) {

	//var nameSpace string = "meta"

	//var path_ip_detect string = fmt.Sprintf("/%s/%s/v1/errorcode", prefix, nameSpace)

	//r.Handle(path_ip_detect, middleware.Adapt(rangeSetHandler(), middleware.Notify()))

}
