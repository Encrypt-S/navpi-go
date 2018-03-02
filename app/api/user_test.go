package api

import (
	"testing"
)

func TestHashDetails(t *testing.T) {

	username := "user"
	password := "pass"

	hashStr, err := HashDetails(username, password)
	if err != nil {
		t.Fatalf("HashDetails error: %s", err)
	}

	if hashStr == "" {
		t.Error("HashDetails HashStr is empty")
	}
}

func TestCheckHashDetails(t *testing.T) {

	username := "user"
	password := "pass"

	hashStr, err := HashDetails(username, password)
	if err != nil {
		t.Fatalf("HashDetails error: %s", err)
	}

	bool := CheckHashDetails(username, password, hashStr)
	if bool != true {
		t.Error("CheckHashDetails should have returned true")
	}

}
