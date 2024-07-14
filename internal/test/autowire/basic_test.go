// Package autowire testing all possible annotation scenarios, multiple dependencies, multiple resolution and circular dependency handling
package autowiretest

import (
	"testing"

	godi "github.com/olbrichattila/godi"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func TestRunner(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type exampleImpl struct {
	Dependency1 exampleDepInterface `di:"autowire"`
}

type exampleDepInterface interface {
	mockFunc()
	mockFuncCalled() int
}

type exampleDep struct {
	mockCalls int
}

func (t *exampleDep) mockFunc() {
	t.mockCalls++
}

func (t *exampleDep) mockFuncCalled() int {
	return t.mockCalls
}

func (t *TestSuite) TestBasicFunctionality() {
	container := godi.New()
	container.Set("olbrichattila.godi.internal.test.autowire.exampleDepInterface", &exampleDep{})
	exampleImp := &exampleImpl{}

	_, err := container.Get(exampleImp)
	t.Nil(err)

	t.Equal(0, exampleImp.Dependency1.mockFuncCalled())
	exampleImp.Dependency1.mockFunc()
	exampleImp.Dependency1.mockFunc()
	t.Equal(2, exampleImp.Dependency1.mockFuncCalled())
}
