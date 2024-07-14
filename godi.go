// Package gody is the wrapper around dependency injection container
package godi

import "github.com/olbrichattila/godi/internal/container"

// Container is an just a wrapper around the internal container to be able to modify the implementation without effecting end users if they would think of referencing the sub package
type Container interface {
	Set(string, interface{})
	Get(interface{}) (interface{}, error)
	Build(map[string]interface{})
	GetDependency(string) (interface{}, error)
	GetDependencies() map[string]interface{}
	Flush()
	Delete(paramName string)
	Count() int
}

// Cont is the container structure which follows Container interface
type Cont struct {
	c *container.Cont
}

// New creates a new container
func New() Container {
	return &Cont{c: container.New()}
}

// Build entire dependency tree
func (t *Cont) Build(dependencies map[string]interface{}) {
	t.c.Build(dependencies)
}

// Set Set dependencies to the container
func (t *Cont) Set(paramName string, dependency interface{}) {
	t.c.Set(paramName, dependency)
}

// Get resolves dependencies. Use a construct function with your dependency interface type hints. They will be resolved recursively.
func (t *Cont) Get(obj interface{}) (interface{}, error) {
	return t.c.Get(obj)
}

// GetDependency retrieve the dependency, or returns error
func (t *Cont) GetDependency(paramName string) (interface{}, error) {
	return t.c.GetDependency(paramName)
}

// GetDependencies returns the entire dependency map
func (t *Cont) GetDependencies() map[string]interface{} {
	return t.c.GetDependencies()
}

// Flush dependencies
func (t *Cont) Flush() {
	t.c.Flush()
}

// Delete one dependency
func (t *Cont) Delete(paramName string) {
	t.c.Delete(paramName)
}

// Count returns how any dependencies provided
func (t *Cont) Count() int {
	return t.c.Count()
}
