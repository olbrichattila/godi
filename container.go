// Package container implements simple Dependency injection container
package godi

import (
	"errors"
	"reflect"
	"strings"
)

var (
	errCannotBeResolved = errors.New("the DI parameter cannot be resolved")
)

// New creates new dependency injector container
func New() *cont {
	c := &cont{}
	c.dependencies = make(map[string]interface{})
	return c
}

type cont struct {
	dependencies map[string]interface{}
}

// Set new dependency, provide a "packagepath.Interfacename" as a string, and your dependency, which always should be an interface or struct
func (t *cont) Set(paramName string, dependency interface{}) {
	t.dependencies[paramName] = dependency
}

// Get will resolve dependencies. Use a Construct func with your dependency interface type hints. They will be resolved recursively
func (t *cont) Get(obj interface{}) (interface{}, error) {
	v := reflect.ValueOf(obj)

	if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct {
		rt := reflect.TypeOf(obj)

		method, found := rt.MethodByName("Construct")
		if found {
			var passParams []reflect.Value
			passParams = append(passParams, v)
			methodType := method.Type
			numParams := methodType.NumIn()

			for i := 1; i < numParams; i++ {
				paramType := methodType.In(i)
				param, err := t.resolve(paramType)
				if err != nil {
					return nil, err
				}

				_, err = t.Get(param)
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

func (t *cont) resolve(paramType reflect.Type) (interface{}, error) {
	pkgPath := paramType.PkgPath() + "/" + paramType.Name()
	fullTypeName := strings.Join(strings.Split(pkgPath, "/")[1:], ".")
	param, ok := t.dependencies[fullTypeName]
	if !ok {
		return nil, errCannotBeResolved
	}

	return param, nil
}
