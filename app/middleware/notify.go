package middleware

import (
	"log"
	"net/http"
)

func Notify() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("before")
			defer log.Println("after")
			h.ServeHTTP(w, r)
		})
	}
}
