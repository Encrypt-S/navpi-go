package setupapi

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type SetupApiSuite struct {
	suite.Suite
}

func (suite *SetupApiSuite) InitSetupHandlersSuccess() {

	suite.Equal(expectedURL, result)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(SetupApiSuite))
}
