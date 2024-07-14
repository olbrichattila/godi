// Package autowire testing all possible annotation scenarios, multiple dependencies, multiple resolution and circular dependency handling
package autowiretest

import (
	"testing"

	godi "github.com/olbrichattila/godi"
	"github.com/stretchr/testify/suite"
)

type TestSuiteMultiple struct {
	suite.Suite
}

func TestRunnerMultiple(t *testing.T) {
	suite.Run(t, new(TestSuiteMultiple))
}

type exampleMultipleImpl struct {
	Dependency1 exampleMultipleDepInterface `di:"autowire"`
	Dependency2 exampleMultipleDepInterface `di:"autowire"`
	Dependency3 exampleMultipleDepInterface `di:"autowire"`
}

type exampleMultipleDepInterface interface {
	mockFunc()
	mockFuncCalled() int
}

type exampleMultipleDep struct {
	mockCalls int
}

func (t *exampleMultipleDep) mockFunc() {
	t.mockCalls++
}

func (t *exampleMultipleDep) mockFuncCalled() int {
	return t.mockCalls
}

func (t *TestSuiteMultiple) TestBasicFunctionality() {
	container := godi.New()
	container.Set("olbrichattila.godi.internal.autowire-test.exampleMultipleDepInterface", &exampleMultipleDep{})
	exampleImp := &exampleMultipleImpl{}

	_, err := container.Get(exampleImp)

	t.Nil(err)

	t.Equal(0, exampleImp.Dependency1.mockFuncCalled())
	t.Equal(0, exampleImp.Dependency2.mockFuncCalled())
	t.Equal(0, exampleImp.Dependency3.mockFuncCalled())

	exampleImp.Dependency1.mockFunc()

	exampleImp.Dependency2.mockFunc()
	exampleImp.Dependency2.mockFunc()

	exampleImp.Dependency3.mockFunc()
	exampleImp.Dependency3.mockFunc()
	exampleImp.Dependency3.mockFunc()

	// They all share the same dependency, calls should be added together
	t.Equal(6, exampleImp.Dependency1.mockFuncCalled())
	t.Equal(6, exampleImp.Dependency2.mockFuncCalled())
	t.Equal(6, exampleImp.Dependency3.mockFuncCalled())
}
