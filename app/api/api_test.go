package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

// metaErrorDisplayHandler test
func Test_metaErrorDisplayHandler(t *testing.T) {
	r := gofight.New()

	r.GET("/").
		SetDebug(true).
		Run(metaErrorDisplayHandler(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			appResp := Response{}
			appResp.Data = AppRespErrors
			jsonBytes, _ := json.Marshal(appResp)
			jsonStr := fmt.Sprintf("%s", jsonBytes)

			assert.Equal(t, jsonStr, r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)

		})
}

func Test_RouteBuilder(t *testing.T) {

	route := RouteBuilder("api", "chain", "v1", "getuser")

	assert.Equal(t, "/api/chain/v1/getuser", route)

}
