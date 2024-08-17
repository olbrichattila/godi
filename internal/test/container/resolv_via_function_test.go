// Package containertest testing all possible scenarios, multiple dependencies, multiple resolution and circular dependency handling
package containertest

import (
	"testing"

	godi "github.com/olbrichattila/godi"
	"github.com/stretchr/testify/suite"
)

type TestFunctionSuite struct {
	suite.Suite
}

func TestFunctionRunner(t *testing.T) {
	suite.Run(t, new(TestFunctionSuite))
}

func (t *TestFunctionSuite) TestBasicFunctionality() {
	container := godi.New()

	oneParamConstructorMock := newOneParamConstructorMock()

	container.Set("olbrichattila.godi.internal.test.container.noParamConstructorInterface", callerFunc)

	_, err := container.Get(oneParamConstructorMock)
	t.Nil(err)

	t.Equal(1, oneParamConstructorMock.ConstructorCallCount())
}

func callerFunc() interface{} {
	return newNoParamConstructorMock()
}
