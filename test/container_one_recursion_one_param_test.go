package test

import (
	"fmt"
	"testing"

	godi "github.com/olbrichattila/godi/internal"
	"github.com/stretchr/testify/suite"
)

type OneRecursionTestSuite struct {
	suite.Suite
}

func TestRunnerOneRecursion(t *testing.T) {
	suite.Run(t, new(OneRecursionTestSuite))
}

func (t *OneRecursionTestSuite) TestFirstRecursionReturnsErrorIfMappingNotSet() {
	container := godi.New()
	oneParamConstructorMock := newOneParamConstructorMock()

	_, err := container.Get(oneParamConstructorMock)
	if err != nil {
		fmt.Println(err.Error())
	}
	t.ErrorIs(err, godi.ErrCannotBeResolved)

}

func (t *OneRecursionTestSuite) TestFirstRecursionResolveSecond() {
	container := godi.New()

	noParamConstructorMock := newNoParamConstructorMock()
	oneParamConstructorMock := newOneParamConstructorMock()

	container.Set("olbrichattila.godi.test.noParamConstructorInterface", noParamConstructorMock)

	_, err := container.Get(oneParamConstructorMock)
	t.Nil(err)

	t.Equal(1, oneParamConstructorMock.ConstructorCallCount())
	t.Equal(1, noParamConstructorMock.ConstructorCallCount())
}

func (t *OneRecursionTestSuite) TestFirstRecursionResolveMultiple() {
	container := godi.New()

	noParamConstructorMock := newNoParamConstructorMock()
	oneParamConstructorMock := newOneParamConstructorMock()

	container.Set("olbrichattila.godi.test.noParamConstructorInterface", noParamConstructorMock)

	_, err := container.Get(oneParamConstructorMock)
	t.Nil(err)

	t.Equal(1, oneParamConstructorMock.ConstructorCallCount())
	t.Equal(1, noParamConstructorMock.ConstructorCallCount())
}
