package user

import (
	"testing"
	"github.com/appleboy/gofight"
	"github.com/NAVCoin/navpi-go/app/conf"
	"github.com/NAVCoin/navpi-go/app/api"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strings"
	"net/http"
)


// metaErrorDisplayHandler test
func Test_loginHandler_correct(t *testing.T) {


	api.BuildAppErrors()
	// setup the tests
	username := "user"
	password := "password"

	hash, _ := api.HashDetails(username, password)
	conf.AppConf.UIPassword = hash


	r := gofight.New()
	r.POST("/").
		SetJSON(gofight.D{
			"username": username,
			"password": password,
		}).
		SetDebug(true).
		Run(loginHandler(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

		var apiResp api.Response
		//
		// get the json from the post data
		err := json.NewDecoder(r.Body).Decode(&apiResp)

		if err != nil {
			t.Error(err.Error())
		}

		assert.Equal(t, r.Code, http.StatusOK)
		assert.NotEqual(t, len(apiResp.Data.(string)), 0)

		//base test we have a jwt
		split := strings.Split(apiResp.Data.(string), ".")
		assert.NotEqual(t, len(split), 2)

	})

}



func Test_loginHandler_no_post_data(t *testing.T) {

	api.BuildAppErrors()

	r := gofight.New()
	r.POST("/").
		SetJSON(gofight.D{}).
		SetDebug(true).

		Run(loginHandler(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

		var apiResp api.Response
		//
		// get the json from the post data
		err := json.NewDecoder(r.Body).Decode(&apiResp)

		if err != nil {
			t.Error(err.Error())
		}

		// should be a 400
		assert.Equal(t, r.Code, http.StatusBadRequest)

		// should have 1 error
		assert.Equal(t, len(apiResp.Errors), 1)

		// check it has the right login error
		resErr := apiResp.Errors[0]
		assert.Equal(t, resErr.Code, api.AppRespErrors.LoginError.Code)

	})

}


func Test_loginHandler_wrong_username(t *testing.T) {


	api.BuildAppErrors()
	// setup the tests
	username := "user"
	password := "password"

	hash, _ := api.HashDetails(username, password)
	conf.AppConf.UIPassword = hash


	r := gofight.New()
	r.POST("/").
		SetJSON(gofight.D{
			"username": "wrong_username",
			"password": password,
		}).
		SetDebug(true).

		Run(loginHandler(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

		var apiResp api.Response
		//
		// get the json from the post data
		err := json.NewDecoder(r.Body).Decode(&apiResp)

		if err != nil {
			t.Error(err.Error())
		}

		// should be a 400
		assert.Equal(t, r.Code, http.StatusBadRequest)

		// should have 1 error
		assert.Equal(t, len(apiResp.Errors), 1)

		// check it has the right login error
		resErr := apiResp.Errors[0]
		assert.Equal(t, resErr.Code, api.AppRespErrors.LoginError.Code)

	})

}



func Test_loginHandler_wrong_password(t *testing.T) {


	api.BuildAppErrors()
	// setup the tests
	username := "user"
	password := "password"

	hash, _ := api.HashDetails(username, password)
	conf.AppConf.UIPassword = hash



	r := gofight.New()
	r.POST("/").
		SetJSON(gofight.D{
		"username": username,
		"password": "wrong_password",
	}).
		SetDebug(true).

		Run(loginHandler(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

		var apiResp api.Response
		//
		// get the json from the post data
		err := json.NewDecoder(r.Body).Decode(&apiResp)

		if err != nil {
			t.Error(err.Error())
		}

		// should be a 400
		assert.Equal(t, r.Code, http.StatusBadRequest)

		// should have 1 error
		assert.Equal(t, len(apiResp.Errors), 1)

		// check it has the right login error
		resErr := apiResp.Errors[0]
		assert.Equal(t, resErr.Code, api.AppRespErrors.LoginError.Code)

	})

}