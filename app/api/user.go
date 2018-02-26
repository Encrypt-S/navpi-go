package api

import "golang.org/x/crypto/bcrypt"

func HashDetails(usernme string, password string) (string, error) {

	return hashStr(username + ":" + password)

}



func hashStr(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

