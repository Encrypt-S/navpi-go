package middleware

import (
	"net/http"
)

// CORSHandler allows http to be served
// adding necessary headers to response
func CORSHandler() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Add("Access-Control-Allow-Headers", "*")
			w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

			// if this is the preflight then exit here
			if r.Method == http.MethodOptions {

				w.WriteHeader(http.StatusOK)
				w.Write([]byte{})
				return

			}

			// all good continue
			h.ServeHTTP(w, r)

		})
	}
}
