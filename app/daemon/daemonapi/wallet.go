package daemonapi

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"encoding/json"

	"github.com/Encrypt-S/navpi-go/app/api"
	"github.com/Encrypt-S/navpi-go/app/conf"
	"github.com/Encrypt-S/navpi-go/app/daemon/daemonrpc"
	"github.com/Encrypt-S/navpi-go/app/middleware"
	"github.com/gorilla/mux"
	"github.com/muesli/crunchy"
)

// InitWalletHandlers sets up handlers for the blockchain rpc interface
func InitWalletHandlers(r *mux.Router, prefix string) {

	namespace := "wallet"

	// setup getstakereport
	stakeReportPath := api.RouteBuilder(prefix, namespace, "v1", "stakeReport")
	r.Handle(stakeReportPath, middleware.Adapt(stakeReport()))

	// setup encryptwallet
	r.Handle(api.RouteBuilder(prefix, namespace, "v1", "encryptwallet"),
		middleware.Adapt(encryptWallet())).
			Methods("POST")

}

// checkPasswordStrength ensures password entered is safe
func checkPasswordStrength(pass string) error {

	validator := crunchy.NewValidator()
	err := validator.Check(pass)

	return err

}

// EncryptWalletCmd defines the "encryptwallet" JSON-RPC command.
type EncryptWalletCmd struct {
	PassPhrase string `json:"passPhrase"`
}

// encryptWallet executes "encryptwallet" json RPC command.
func encryptWallet() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var encryptWalletCmd EncryptWalletCmd
		apiResp := api.Response{}

		err := json.NewDecoder(r.Body).Decode(&encryptWalletCmd)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			returnErr := api.AppRespErrors.ServerError
			returnErr.ErrorMessage = fmt.Sprintf("Server error: %v", err)
			apiResp.Errors = append(apiResp.Errors, returnErr)
			apiResp.Send(w)

			return
		}

		err = checkPasswordStrength(encryptWalletCmd.PassPhrase)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			returnErr := api.AppRespErrors.InvalidStrength
			returnErr.ErrorMessage = fmt.Sprintf("Invalid strength error: %v", err)
			apiResp.Errors = append(apiResp.Errors, returnErr)
			apiResp.Send(w)
			return
		}

		n := daemonrpc.RpcRequestData{}
		n.Method = "encryptwallet"
		n.Params = []string{encryptWalletCmd.PassPhrase}

		resp, err := daemonrpc.RequestDaemon(n, conf.NavConf)

		if err != nil {
			daemonrpc.RpcFailed(err, w, r)
			return
		}

		bodyText, err := ioutil.ReadAll(resp.Body)
		w.WriteHeader(resp.StatusCode)
		w.Write(bodyText)

	})
}

// getstakereport: return SubTotal of the staked coin in last 24H, 7 days, etc.. of all owns address
func stakeReport() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		n := daemonrpc.RpcRequestData{}
		n.Method = "getstakereport"
		//
		resp, err := daemonrpc.RequestDaemon(n, conf.NavConf)

			// Handle errors requesting the daemon
			if err != nil {
				daemonrpc.RpcFailed(err, w, r)
				return
			}

		bodyText, err := ioutil.ReadAll(resp.Body)
		w.WriteHeader(resp.StatusCode)
		w.Write(bodyText)

		return

	})
}
