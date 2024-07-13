// Package godi is a golang dependency injector container
package godi

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	// ErrCannotBeResolved returned when the container not able to resolve the dependency, not mapped with container.Set()
	ErrCannotBeResolved = errors.New("the DI parameter cannot be resolved")
	// ErrCircularReference returned when the dependencies would end up in a forever loop. instead golang blowing up, it returns an error.
	ErrCircularReference = errors.New("circular reference")
)

type dependencyMap struct {
	dependency interface{}
}

// New creates a new dependency injector container
func New() *Cont {
	c := &Cont{}
	c.dependencies = make(map[string]*dependencyMap)
	return c
}

// Cont is the container returned by New
type Cont struct {
	dependencies map[string]*dependencyMap
}

// Set new dependency, provide a "packagePath.InterfaceName" as a string, and your dependency, which should always be an interface or struct
func (t *Cont) Set(paramName string, dependency interface{}) {
	t.dependencies[paramName] = &dependencyMap{dependency: dependency}
}

// Get resolves dependencies. Use a Construct func with your dependency interface type hints. They will be resolved recursively
func (t *Cont) Get(obj interface{}) (interface{}, error) {
	callStack := make(map[string]bool)
	return t.getRecursive(obj, callStack)
}

// getRecursive resolves dependencies recursively, tracking call stack to detect circular references
func (t *Cont) getRecursive(obj interface{}, callStack map[string]bool) (interface{}, error) {
	v := reflect.ValueOf(obj)

	if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct {
		rt := reflect.TypeOf(obj)

		method, found := rt.MethodByName("Construct")
		if found {
			passParams := []reflect.Value{v}
			methodType := method.Type
			numParams := methodType.NumIn()

			for i := 1; i < numParams; i++ {
				paramType := methodType.In(i)
				param, fullTypeName, err := t.resolve(paramType)
				if err != nil {
					return nil, err
				}
				if callStack[fullTypeName] {
					return nil, errors.Join(ErrCircularReference, fmt.Errorf("circular call: %s", fullTypeName))
				}
				callStack[fullTypeName] = true

				_, err = t.getRecursive(param, callStack)
				delete(callStack, fullTypeName)
				if err != nil {
					return nil, err
				}
				passParams = append(passParams, reflect.ValueOf(param))
			}

			method.Func.Call(passParams)
		}
	}
	return obj, nil
}

// resolve finds and returns the dependency, checking for circular references
func (t *Cont) resolve(paramType reflect.Type) (interface{}, string, error) {
	pkgPath := paramType.PkgPath() + "/" + paramType.Name()
	fullTypeName := strings.Join(strings.Split(pkgPath, "/")[1:], ".")
	param, ok := t.dependencies[fullTypeName]
	if !ok {
		return nil, fullTypeName, errors.Join(ErrCannotBeResolved, fmt.Errorf("dependency name: %s", fullTypeName))
	}

	return param.dependency, fullTypeName, nil
}