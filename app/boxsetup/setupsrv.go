package boxsetup

import (
	"github.com/NAVCoin/navpi-go/app/conf"
	"net/http"
	"fmt"
	"io"
)

var setupServer *http.Server

func StartServer (serverConfig conf.ServerConfig) *http.Server {

	port := fmt.Sprintf(":%d", serverConfig.DaemonApiPort)

	srv := &http.Server{Addr: port}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})

	go func() {
		srv.ListenAndServe()
		//http.ListenAndServe("localhost:8081", serverMuxA)
	}()

	// store it so we can get it later
	setupServer = srv
	return srv
}
