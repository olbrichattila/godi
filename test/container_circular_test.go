package test

import (
	"testing"

	godi "github.com/olbrichattila/godi/internal"
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

	container.Set("olbrichattila.godi.test.nestedCircularInterface", nestedCircularMock)

	_, err := container.Get(nestedCircularMock)
	t.Error(err)
	t.ErrorIs(err, godi.ErrCircularReference)
}

func (t *CircularTestSuite) TestMultpleCircularNesting() {
	container := godi.New()

	nestedCircularFirstMock := newNestedCircularFirstMock()
	nestedCircularSecondMock := newNestedCircularSecondMock()
	nestedCircularCircularMock := newNestedCircularThirdMock()

	container.Set("olbrichattila.godi.test.nestedCircularFirstInterface", nestedCircularFirstMock)
	container.Set("olbrichattila.godi.test.nestedCircularSecondInterface", nestedCircularSecondMock)
	container.Set("olbrichattila.godi.test.nestedCircularThirdInterface", nestedCircularCircularMock)

	_, err := container.Get(nestedCircularFirstMock)
	t.Error(err)
	t.ErrorIs(err, godi.ErrCircularReference)
}
