package api

import (
	"testing"
)

func TestHashDetails(t *testing.T) {

	username := "user"
	password := "pass"

	hashdetails, err := HashDetails(username, password)

	if err != nil {
		t.Error("error", err)
	}

	if hashdetails != "" {
		t.Log("success", hashdetails)
	}

}
