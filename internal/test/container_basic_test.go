// Package test testing all possible scenarios, multiple dependencies, multiple resolution and circular dependency handling
package test

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

type exampleImpl struct{}

func (t *TestSuite) TestBasicFunctionality() {
	container := godi.New()
	container.Set("exampleImpl", &exampleImpl{})

	instance, err := container.Get(&exampleImpl{})
	t.Nil(err)

	_, ok := instance.(*exampleImpl)
	t.True(ok)
}
