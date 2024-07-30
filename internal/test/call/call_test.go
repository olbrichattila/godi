// Package autowire testing all possible annotation scenarios, multiple dependencies, multiple resolution and circular dependency handling
package calltest

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

type exampleDepInterface interface {
	mockFunc()
	mockPars(string, int)
	mockFuncCalled() int
	getMockPars() (string, int)
}

type exampleDep struct {
	mockCalls  int
	mockString string
	mockInt    int
}

func (t *exampleDep) mockFunc() {
	t.mockCalls++
}

func (t *exampleDep) mockPars(s string, i int) {
	t.mockString = s
	t.mockInt = i
}

func (t *exampleDep) getMockPars() (string, int) {
	return t.mockString, t.mockInt
}

func (t *exampleDep) mockFuncCalled() int {
	return t.mockCalls
}

func exampleFunc(e exampleDepInterface) {
	e.mockFunc()
	e.mockFunc()
}

func exampleFuncWithExtraParams(par1 string, par2 int, e exampleDepInterface) {
	e.mockFunc()
	e.mockFunc()
	e.mockPars(par1, par2)
}

func (t *TestSuite) TestBasicFunctionality() {
	container := godi.New()
	dependency := &exampleDep{}
	container.Set("olbrichattila.godi.internal.test.call.exampleDepInterface", dependency)

	_, err := container.Call(exampleFunc)
	t.Nil(err)
	t.Equal(2, dependency.mockFuncCalled())
}

func (t *TestSuite) TestWithExtraParameters() {
	container := godi.New()
	dependency := &exampleDep{}
	container.Set("olbrichattila.godi.internal.test.call.exampleDepInterface", dependency)

	_, err := container.Call(exampleFuncWithExtraParams, "test", 500)
	t.Nil(err)
	t.Equal(2, dependency.mockFuncCalled())

	mockString, mockInt := dependency.getMockPars()
	t.Equal("test", mockString)
	t.Equal(500, mockInt)
}
