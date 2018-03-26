package daemonapi

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"encoding/json"

	"github.com/NAVCoin/navpi-go/app/api"
	"github.com/NAVCoin/navpi-go/app/conf"
	"github.com/NAVCoin/navpi-go/app/daemon/daemonrpc"
	"github.com/gorilla/mux"
	"github.com/muesli/crunchy"
)

// InitWalletHandlers sets up handlers for the blockchain rpc interface
func InitWalletHandlers(r *mux.Router, prefix string) {

	namespace := "wallet"
	r.HandleFunc(fmt.Sprintf("/%s/%s/v1/getstakereport", prefix, namespace), getStakeReport).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/%s/v1/encryptwallet", prefix, namespace), encryptWallet).Methods("POST")

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

// getStakeReport takes writer, request - writes out stake report
func getStakeReport(w http.ResponseWriter, r *http.Request) {

	n := daemonrpc.RpcRequestData{}
	n.Method = "getstakereport"

	resp, err := daemonrpc.RequestDaemon(n, conf.NavConf)

	// Handle errors requesting the daemon
	if err != nil {
		daemonrpc.RpcFailed(err, w, r)
		return
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Write(bodyText)
}
