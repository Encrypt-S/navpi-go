package setupapi

import (
	"testing"
	"github.com/stretchr/testify/suite"
)

// setup the test suite
type SetupApiSuite struct {
	suite.Suite
}

// InitSetupHandlers Test
func (suite *SetupApiSuite) InitSetupHandlersSuccess() {

	suite.Equal(expectedURL, result)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(SetupApiSuite))
}