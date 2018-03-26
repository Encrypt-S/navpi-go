package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// GenerateJWT accepts a duration of
// time that will be added to the current
// time to generate the expiry date
func GenerateJWT(exp time.Duration, signingKey []byte) string {

	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	//claims["admin"] = true
	claims["exp"] = time.Now().Add(exp).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(signingKey)

	return tokenString

}
