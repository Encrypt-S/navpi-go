package setupapi

import (
	"testing"
	"github.com/stretchr/testify/suite"
)

package gobotto

import (
"github.com/stretchr/testify/suite"
"testing"
)

// setup the test suite
type SetupApiSuite struct {
	suite.Suite
}

// InitSetupHandlers Test
func (suite *SetupApiSuite) InitSetupHandlersSuccess() {
	expectedURL := "http://my-cool-domain.com/robots.txt"
	result, _ := RobotsURL("http://my-cool-domain.com/blog-post/1")

	// Notice we are now using `suite` to call the assertion methods
	suite.Equal(expectedURL, result)
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(SetupApiSuite))
}