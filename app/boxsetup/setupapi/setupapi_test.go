package setupapi

import (
	"fmt"
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"github.com/gorilla/mux"
	"log"
)

// testing prototype :: gofight API Handler

func TestInitSetupHandlers(t *testing.T) {
	r := gofight.New()

	var nameSpace string = "setup"
	var prefix string = "api"

	router := mux.NewRouter()

	r.GET(fmt.Sprintf("/%s/%s/v1/setrange", prefix, nameSpace)).
		// turn on the debug mode.
		SetDebug(true).
		Run(InitSetupHandlers(router, prefix), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			log.Println(r.Body.String())
			log.Println(r.Code)

			assert.Equal(t, "Hello World", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

//func TestInitSetupHandlers(t *testing.T) {
//
//	var nameSpace string = "setup"
//
//	req, err := http.NewRequest("GET", "/%s/%s/v1/setrange", nameSpace)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(HealthCheckHandler)
//
//	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
//	// directly and pass in our Request and ResponseRecorder.
//	handler.ServeHTTP(rr, req)
//
//	// Check the status code is what we expect.
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//
//	// Check the response body is what we expect.
//	expected := `{"alive": true}`
//	if rr.Body.String() != expected {
//		t.Errorf("handler returned unexpected body: got %v want %v",
//			rr.Body.String(), expected)
//	}
//}
