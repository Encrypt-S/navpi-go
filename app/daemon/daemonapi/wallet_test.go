package daemonapi

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Encrypt-S/navpi-go/app/api"
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
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
	}
	validPws = []string{"d1924ce3d0510b2b2b4604c99453e2e1"}
)

func Test_checkPasswordStrength(t *testing.T) {

	// run through valid password range
	for _, pw := range validPws {
		err := checkPasswordStrength(pw)
		if err != nil {
			t.Errorf("Expected no error for valid password '%s', got %v", pw, err)
		}
	}

	// run through invalid password range
	for _, pw := range invalidPws {
		err := checkPasswordStrength(pw)
		if err == nil {
			t.Errorf("Expected error for invalid password '%s',  got %v", pw, err)
		}
	}

}

// Test encryptWallet func with and empty passphrase
func Test_encryptWallet_empty_passphrase(t *testing.T) {

	api.BuildAppErrors()

	r := gofight.New()
	r.POST("/").
		SetJSON(gofight.D{
			"passPhrase": "",
		}).
		SetDebug(true).
		Run(encryptWallet(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			var apiResp api.Response

			// get the json from the post data
			err := json.NewDecoder(r.Body).Decode(&apiResp)

			if err != nil {
				t.Error(err.Error())
			}

			assert.Equal(t, r.Code, http.StatusBadRequest)
			assert.NotNil(t, apiResp.Errors)
			assert.Equal(t, len(apiResp.Errors), 1)

		})

}

// Test encryptWallet func without a passphrase
func Test_encryptWallet_no_passphrase(t *testing.T) {

	api.BuildAppErrors()

	r := gofight.New()
	r.POST("/").
		SetDebug(true).
		Run(encryptWallet(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			var apiResp api.Response

			// get the json from the post data
			err := json.NewDecoder(r.Body).Decode(&apiResp)

			if err != nil {
				t.Error(err.Error())
			}

			assert.Equal(t, r.Code, http.StatusInternalServerError)
			assert.NotNil(t, apiResp.Errors)
			assert.Equal(t, len(apiResp.Errors), 1)

		})

}

// Test getStakeReport func
func Test_encryptWallets(t *testing.T) {

	api.BuildAppErrors()

	r := gofight.New()
	r.GET("/").
		SetDebug(true).
		Run(stakeReport(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			var apiResp api.Response

			// get the json from the response body
			err := json.NewDecoder(r.Body).Decode(&apiResp)

			if err != nil {
				t.Error(err.Error())
			}

			// ensure we have success
			assert.Nil(t, apiResp.Errors)
			assert.Equal(t, http.StatusOK, r.Code)

		})

}
