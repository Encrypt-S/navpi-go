package middleware

import (
	"net/http"
)

func CorsHandler() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Add("Access-Control-Allow-Headers", "*")

			// if this is the preflight then exit heres
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
