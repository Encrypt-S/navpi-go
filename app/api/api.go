package api

import (
	"github.com/gorilla/mux"
	"fmt"
	"github.com/NAVCoin/navpi-go/app/middleware"
)

type APIResponse struct {

	Data string `json:"data"`
	Success bool `json:"success"`
	Error APIError `json:"error"`

}

type APIError struct {
	Code string `json:"code"`
	Message string `json:"message"`
}


type ErrorCode struct {
	 errorCode string
	 ErrorMessage string
}


//errorCodes := AppErrorCodes{}






func InitMetaHandlers(r *mux.Router, prefix string) {

	var nameSpace string = "meta"

	var path_ip_detect string = fmt.Sprintf("/%s/%s/v1/errorcode", prefix, nameSpace)

	r.Handle(path_ip_detect, middleware.Adapt(rangeSetHandler(), middleware.Notify()))

}
