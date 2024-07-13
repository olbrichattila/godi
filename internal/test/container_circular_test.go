package test

import (
	"testing"

	godi "github.com/olbrichattila/godi/pkg"
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

	container.Set("olbrichattila.godi.internal.test.nestedCircularInterface", nestedCircularMock)

	_, err := container.Get(nestedCircularMock)
	t.Error(err)
	t.ErrorIs(err, godi.ErrCircularReference)
}

func (t *CircularTestSuite) TestMultpleCircularNesting() {
	container := godi.New()

	nestedCircularFirstMock := newNestedCircularFirstMock()
	nestedCircularSecondMock := newNestedCircularSecondMock()
	nestedCircularCircularMock := newNestedCircularThirdMock()

	container.Set("olbrichattila.godi.internal.test.nestedCircularFirstInterface", nestedCircularFirstMock)
	container.Set("olbrichattila.godi.internal.test.nestedCircularSecondInterface", nestedCircularSecondMock)
	container.Set("olbrichattila.godi.internal.test.nestedCircularThirdInterface", nestedCircularCircularMock)

	_, err := container.Get(nestedCircularFirstMock)
	t.Error(err)
	t.ErrorIs(err, godi.ErrCircularReference)
}
