package api

import "golang.org/x/crypto/bcrypt"

// HashDetails takes username and password - returns hash
func HashDetails(username string, password string) (string, error) {
	return hashStr(username + ":" + password)
}

// CheckHashDetails takes username, password, hash - returns checked version
func CheckHashDetails(username string, password string, hash string) bool {
	return checkHashStr(username+":"+password, hash)
}

// hashStr takes password - returns bytes
func hashStr(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checkHashStr takes password, hash - returns true or false
func checkHashStr(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
