package api

import "golang.org/x/crypto/bcrypt"

func HashDetails(username string, password string) (string, error) {
	return hashStr(username + ":" + password)
}

func CheckHashDetails(username string, password string, hash string) bool {
	return checkHashStr(username+":"+password, hash)
}

func hashStr(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkHashStr(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
