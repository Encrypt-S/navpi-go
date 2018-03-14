package daemonapi

import (
	"testing"
)

var (
	invalidPws = []string{
		"",
		" ",
		"aaaaaaaa",
		"crunchy",
		"aaaaaaaa",
		"aabbccdd",
		"12345678",
		"87654321",
		"abcdefgh",
		"hgfedcba",

		"password",
		"intoxicate",
		"p@ssw0rd",    // dictionary with mangling
		"!pass@word?", // dictionary with mangling
		"drowssap",    // reversed dictionary
		"?drow@ssap!", // reversed dictionary with mangling
	}
	validPws = []string{"d1924ce3d0510b2b2b4604c99453e2e1"}
)

func Test_checkPasswordStrength(t *testing.T) {

	for _, pw := range validPws {
		err := checkPasswordStrength(pw)
		if err != nil {
			t.Errorf("Expected no error for valid password '%s', got %v", pw, err)
		}
	}

	for _, pw := range invalidPws {
		err := checkPasswordStrength(pw)
		if err == nil {
			t.Errorf("Expected error for invalid password '%s',  got %v", pw, err)
		}
	}

}
