package middleware

import (
	"net/http"
	"strings"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/NAVCoin/navpi-go/app/conf"
	"fmt"
)

func JwtHandler() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			headerToken, _ := FromAuthHeader(r)

			token, _ := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte(conf.ServerConf.JWTSecret), nil
			})

			if !token.Valid {

				status := http.StatusUnauthorized
				http.Error(w, http.StatusText(status), status)
				return

			}


			// all good continue
			h.ServeHTTP(w, r)
		})
	}
}



// FromAuthHeader is a "TokenExtractor" that takes a give request and extracts
// the JWT token from the Authorization header.
func FromAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.New("Authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}
