package containertest

import (
	"testing"

	godi "github.com/olbrichattila/godi"
	internalcontainer "github.com/olbrichattila/godi/internal/container"
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
	t.ErrorIs(err, internalcontainer.ErrCannotBeResolved)

}

func (t *OneRecursionTestSuite) TestFirstRecursionResolveSecond() {
	container := godi.New()

	noParamConstructorMock := newNoParamConstructorMock()
	oneParamConstructorMock := newOneParamConstructorMock()

	container.Set("olbrichattila.godi.internal.test.container.noParamConstructorInterface", noParamConstructorMock)

	_, err := container.Get(oneParamConstructorMock)
	t.Nil(err)

	t.Equal(1, oneParamConstructorMock.ConstructorCallCount())
	t.Equal(1, noParamConstructorMock.ConstructorCallCount())
}

func (t *OneRecursionTestSuite) TestFirstRecursionResolveMultiple() {
	container := godi.New()

	noParamConstructorMock := newNoParamConstructorMock()
	oneParamConstructorMock := newOneParamConstructorMock()

	container.Set("olbrichattila.godi.internal.test.container.noParamConstructorInterface", noParamConstructorMock)

	_, err := container.Get(oneParamConstructorMock)
	t.Nil(err)

	t.Equal(1, oneParamConstructorMock.ConstructorCallCount())
	t.Equal(1, noParamConstructorMock.ConstructorCallCount())
}
