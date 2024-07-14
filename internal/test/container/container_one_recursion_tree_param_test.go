package containertest

import (
	"testing"

	godi "github.com/olbrichattila/godi"
	"github.com/stretchr/testify/suite"
)

type TreeRecursionTestSuite struct {
	suite.Suite
}

func TestRunnerTreeRecursion(t *testing.T) {
	suite.Run(t, new(TreeRecursionTestSuite))
}

func (t *TreeRecursionTestSuite) TestFirstRecursionResolveMultiple() {
	container := godi.New()

	noParamConstructorMock := newNoParamConstructorMock()
	treeParamConstructorMock := newTreeParamConstructorMock()

	container.Set("olbrichattila.godi.internal.test.container.noParamConstructorInterface", noParamConstructorMock)

	_, err := container.Get(treeParamConstructorMock)
	t.Nil(err)

	t.Equal(1, treeParamConstructorMock.ConstructorCallCount())
	t.Equal(3, noParamConstructorMock.ConstructorCallCount())
}
