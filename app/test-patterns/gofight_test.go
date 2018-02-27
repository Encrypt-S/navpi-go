package main

import (
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// basic test setup :: gofight API Handler

func TestBasicHelloWorld(t *testing.T) {
	r := gofight.New()

	r.GET("/").
		// turn on the debug mode.
		SetDebug(true).
		Run(BasicEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			assert.Equal(t, "Hello World", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
