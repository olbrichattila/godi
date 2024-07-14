package containertest

import (
	"testing"

	godi "github.com/olbrichattila/godi"
	internalcontainer "github.com/olbrichattila/godi/internal/container"
	"github.com/stretchr/testify/suite"
)

type CircularTestSuite struct {
	suite.Suite
}

func TestRunnerCircular(t *testing.T) {
	suite.Run(t, new(CircularTestSuite))
}

func (t *CircularTestSuite) TestDirectCircular() {
	container := godi.New()
	nestedCircularMock := newNestedCircularMock()

	container.Set("olbrichattila.godi.internal.test.container.nestedCircularInterface", nestedCircularMock)

	_, err := container.Get(nestedCircularMock)
	t.Error(err)
	t.ErrorIs(err, internalcontainer.ErrCircularReference)
}

func (t *CircularTestSuite) TestMultpleCircularNesting() {
	container := godi.New()

	nestedCircularFirstMock := newNestedCircularFirstMock()
	nestedCircularSecondMock := newNestedCircularSecondMock()
	nestedCircularCircularMock := newNestedCircularThirdMock()

	container.Set("olbrichattila.godi.internal.test.container.nestedCircularFirstInterface", nestedCircularFirstMock)
	container.Set("olbrichattila.godi.internal.test.container.nestedCircularSecondInterface", nestedCircularSecondMock)
	container.Set("olbrichattila.godi.internal.test.container.nestedCircularThirdInterface", nestedCircularCircularMock)

	_, err := container.Get(nestedCircularFirstMock)
	t.Error(err)
	t.ErrorIs(err, internalcontainer.ErrCircularReference)
}
