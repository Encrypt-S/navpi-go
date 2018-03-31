package middleware

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// defaultAuthorizationHeaderName is the default header name where the Auth
// token should be written
const defaultAuthorizationHeaderName = "Authorization"
const jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o"

// envVarClientSecretName the environment variable to read the JWT environment
// variable
const envVarClientSecretName = "CLIENT_SECRET_VAR_SHHH"

// check that the
func Test_jwt_FromAuthHeader_wrong_format(t *testing.T) {

	// no bearer
	h := http.Header{}
	h.Add(defaultAuthorizationHeaderName, jwtToken)
	r := http.Request{Header: h}

	extractedToken, err := FromAuthHeader(&r)

	assert.Equal(t, "", extractedToken)
	assert.NotNil(t, err)

	// no header - should not error
	h = http.Header{}
	r = http.Request{Header: h}

	extractedToken, err = FromAuthHeader(&r)
	assert.Equal(t, "", extractedToken)
	assert.Nil(t, err)

	// not properly formed auth header
	h = http.Header{}
	h.Add(defaultAuthorizationHeaderName, "bearer")
	r = http.Request{Header: h}

	extractedToken, err = FromAuthHeader(&r)
	assert.Equal(t, "", extractedToken)
	assert.NotNil(t, err)
}

// check that the
func Test_jwt_FromAuthHeader_correct(t *testing.T) {

	// no bearer
	h := http.Header{}
	h.Add(defaultAuthorizationHeaderName, fmt.Sprintf("bearer %s", jwtToken))
	r := http.Request{Header: h}

	extractedToken, err := FromAuthHeader(&r)

	assert.Equal(t, jwtToken, extractedToken)
	assert.Nil(t, err)
}
