// Package autowire testing all possible annotation scenarios, multiple dependencies, multiple resolution and circular dependency handling
package buildertest

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

func (t *TestSuite) TestSet() {
	container := godi.New()
	dependencyName := "testDependency"
	expected := "test1"

	t.Equal(0, container.Count())

	container.Set(dependencyName, expected)
	t.Equal(1, container.Count())
	resolved, err := container.GetDependency(dependencyName)
	t.Nil(err)
	t.Equal(expected, resolved.(string))
}

func (t *TestSuite) TestSetMultiple() {
	container := godi.New()

	t.Equal(0, container.Count())

	dependencyName1 := "testDependency1"
	expected1 := "test1"

	dependencyName2 := "testDependency2"
	expected2 := "test2"

	dependencyName3 := "testDependency3"
	expected3 := "test3"

	container.Set(dependencyName1, expected1)
	container.Set(dependencyName2, expected2)
	container.Set(dependencyName3, expected3)

	t.Equal(3, container.Count())

	resolved, err := container.GetDependency(dependencyName1)
	t.Nil(err)
	t.Equal(expected1, resolved.(string))

	resolved, err = container.GetDependency(dependencyName2)
	t.Nil(err)
	t.Equal(expected2, resolved.(string))

	resolved, err = container.GetDependency(dependencyName3)
	t.Nil(err)
	t.Equal(expected3, resolved.(string))
}

func (t *TestSuite) TestErrorReturnedIfNoDependency() {
	container := godi.New()
	dependencyName := "testDependency"
	expected := "test1"
	notSetDependencyName := "testDependency2"

	t.Equal(0, container.Count())

	container.Set(dependencyName, expected)
	t.Equal(1, container.Count())
	_, err := container.GetDependency(notSetDependencyName)
	t.Error(err)
}

func (t *TestSuite) TestBuildMultiple() {
	container := godi.New()

	t.Equal(0, container.Count())

	dependencyName1 := "testDependency1"
	expected1 := "test1"

	dependencyName2 := "testDependency2"
	expected2 := "test2"

	dependencyName3 := "testDependency3"
	expected3 := "test3"

	dependencies := map[string]interface{}{
		dependencyName1: expected1,
		dependencyName2: expected2,
		dependencyName3: expected3,
	}

	container.Build(dependencies)

	t.Equal(3, container.Count())

	resolved, err := container.GetDependency(dependencyName1)
	t.Nil(err)
	t.Equal(expected1, resolved.(string))

	resolved, err = container.GetDependency(dependencyName2)
	t.Nil(err)
	t.Equal(expected2, resolved.(string))

	resolved, err = container.GetDependency(dependencyName3)
	t.Nil(err)
	t.Equal(expected3, resolved.(string))
}

func (t *TestSuite) TestFlush() {
	container := godi.New()

	t.Equal(0, container.Count())

	dependencies := map[string]interface{}{
		"dependencyName1": "expected1",
		"dependencyName2": "expected2",
		"dependencyName3": "expected3",
	}

	container.Build(dependencies)

	t.Equal(3, container.Count())
	container.Flush()
	t.Equal(0, container.Count())
}

func (t *TestSuite) TestDelete() {
	container := godi.New()

	t.Equal(0, container.Count())

	dependencies := map[string]interface{}{
		"dependencyName1": "expected1",
		"dependencyName2": "expected2",
		"dependencyName3": "expected3",
	}

	container.Build(dependencies)

	t.Equal(3, container.Count())
	container.Delete("dependencyName2")
	t.Equal(2, container.Count())
}

func (t *TestSuite) TestGetDependencyMap() {
	container := godi.New()

	deps := container.GetDependencies()
	t.Len(deps, 0)

	dependencies := map[string]interface{}{
		"dependencyName1": "expected1",
		"dependencyName2": "expected2",
		"dependencyName3": "expected3",
	}

	container.Build(dependencies)

	deps = container.GetDependencies()
	t.Len(deps, 3)
}
