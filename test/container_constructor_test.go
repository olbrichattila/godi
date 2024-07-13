package test

import (
	"testing"

	godi "github.com/olbrichattila/godi/internal"
	"github.com/stretchr/testify/suite"
)

type ConstructorTestSuite struct {
	suite.Suite
}

func TestRunnerConstructor(t *testing.T) {
	suite.Run(t, new(ConstructorTestSuite))
}

func (t *ConstructorTestSuite) TestConstructorCalled() {
	container := godi.New()
	noParamConstructorMock := newNoParamConstructorMock()

	container.Get(noParamConstructorMock)

	t.Equal(1, noParamConstructorMock.ConstructorCallCount())
}