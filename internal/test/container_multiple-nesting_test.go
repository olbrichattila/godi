package test

import (
	"testing"

	godi "github.com/olbrichattila/godi/pkg"
	"github.com/stretchr/testify/suite"
)

type NestingTestSuite struct {
	suite.Suite
}

func TestRunnerNesting(t *testing.T) {
	suite.Run(t, new(NestingTestSuite))
}

func (t *NestingTestSuite) TestMultipleNesting() {
	container := godi.New()

	nestedFirstMock := newNestedFirstMock()
	nestedSecondMock := newNestedSecondMock()
	nestedThirdMock := newNestedThirdMock()

	container.Set("olbrichattila.godi.internal.test.nestedSecondInterface", nestedSecondMock)
	container.Set("olbrichattila.godi.internal.test.nestedThirdInterface", nestedThirdMock)

	_, err := container.Get(nestedFirstMock)
	t.Nil(err)

	t.Equal(1, nestedFirstMock.ConstructorCallCount())
	t.Equal(1, nestedSecondMock.ConstructorCallCount())
	t.Equal(1, nestedThirdMock.ConstructorCallCount())

	t.Equal(0, nestedFirstMock.MockFuncCallCount())
	t.Equal(0, nestedSecondMock.MockFuncCallCount())
	t.Equal(0, nestedThirdMock.MockFuncCallCount())

	// Assert functions are relying on their dependencies and calling each other
	nestedFirstMock.MockFunc()

	t.Equal(1, nestedFirstMock.MockFuncCallCount())
	t.Equal(1, nestedSecondMock.MockFuncCallCount())
	t.Equal(1, nestedThirdMock.MockFuncCallCount())
}
