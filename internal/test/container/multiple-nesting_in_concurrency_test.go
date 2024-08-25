package containertest

import (
	"sync"
	"testing"

	godi "github.com/olbrichattila/godi"
	"github.com/stretchr/testify/suite"
)

type ConcurrentNestingTestSuite struct {
	suite.Suite
}

func TestRunnerConcurrentNesting(t *testing.T) {
	suite.Run(t, new(ConcurrentNestingTestSuite))
}

func (t *ConcurrentNestingTestSuite) TestMultipleNesting() {
	container := godi.New()

	nestedFirstMock := newNestedFirstMock()
	nestedSecondMock := newNestedSecondMock()
	nestedThirdMock := newNestedThirdMock()

	container.Set("olbrichattila.godi.internal.test.container.nestedSecondInterface", nestedSecondMock)
	container.Set("olbrichattila.godi.internal.test.container.nestedThirdInterface", nestedThirdMock)

	var wg sync.WaitGroup
	callCount := 15
	wg.Add(callCount)
	for i := 0; i < callCount; i++ {
		go func() {
			defer wg.Done()
			_, err := container.Get(nestedFirstMock)
			t.Nil(err)
		}()
	}

	wg.Wait()

	t.Equal(15, nestedFirstMock.ConstructorCallCount())
	t.Equal(15, nestedSecondMock.ConstructorCallCount())
	t.Equal(15, nestedThirdMock.ConstructorCallCount())

	t.Equal(0, nestedFirstMock.MockFuncCallCount())
	t.Equal(0, nestedSecondMock.MockFuncCallCount())
	t.Equal(0, nestedThirdMock.MockFuncCallCount())

	// Assert functions are relying on their dependencies and calling each other
	nestedFirstMock.MockFunc()

	t.Equal(1, nestedFirstMock.MockFuncCallCount())
	t.Equal(1, nestedSecondMock.MockFuncCallCount())
	t.Equal(1, nestedThirdMock.MockFuncCallCount())
}
