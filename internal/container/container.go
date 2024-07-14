// Package container is a golang dependency injector container
package container

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	// ErrCannotBeResolved is returned when the container is not able to resolve the dependency.
	ErrCannotBeResolved = errors.New("the DI parameter cannot be resolved")
	// ErrCannotBeResolvedPossibleNeedExport is returned when the container is not able to resolve the dependency, possibly due to an unexported field.
	ErrCannotBeResolvedPossibleNeedExport = errors.New("the DI parameter cannot be resolved, possible unexported field for autowire notation")
	// ErrCircularReference is returned when there is a circular dependency reference.
	ErrCircularReference = errors.New("circular reference")
)

// New creates a new dependency injector container.
func New() *Cont {
	return &Cont{
		dependencies: make(map[string]interface{}),
	}
}

// Cont is the container returned by New.
type Cont struct {
	callStack    map[string]bool
	dependencies map[string]interface{}
}

// Build entire dependency tree
func (t *Cont) Build(dependencies map[string]interface{}) {
	t.dependencies = dependencies
}

// Set registers a new dependency. Provide a "packagePath.InterfaceName" as a string and your dependency, which should be an interface or struct.
func (t *Cont) Set(paramName string, dependency interface{}) {
	t.dependencies[paramName] = dependency
}

// GetDependency retrieve the dependency, or returns error
func (t *Cont) GetDependency(paramName string) (interface{}, error) {
	if dep, ok := t.dependencies[paramName]; ok {
		return dep, nil
	}

	return nil, fmt.Errorf("cannot retrieve dependency %s", paramName)
}

// GetDependencies returns the entire dependency map
func (t *Cont) GetDependencies() map[string]interface{} {
	return t.dependencies
}

// Flush dependencies
func (t *Cont) Flush() {
	t.dependencies = make(map[string]interface{})
}

// Delete one dependency
func (t *Cont) Delete(paramName string) {
	delete(t.dependencies, paramName)
}

// Count returns how any dependencies provided
func (t *Cont) Count() int {
	return len(t.dependencies)
}

// Get resolves dependencies. Use a construct function with your dependency interface type hints. They will be resolved recursively.
func (t *Cont) Get(obj interface{}) (interface{}, error) {
	t.callStack = make(map[string]bool)
	return t.getRecursive(obj)
}

// getRecursive resolves dependencies recursively, tracking call stack to detect circular references.
func (t *Cont) getRecursive(obj interface{}) (interface{}, error) {
	v := reflect.ValueOf(obj)

	if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct {
		rt := v.Type()

		if err := t.resolveConstructor(v, rt); err != nil {
			return nil, err
		}

		if err := t.resolveAutoWire(v); err != nil {
			return nil, err
		}
	}

	return obj, nil
}

// resolveConstructor resolves dependencies for the constructor method.
func (t *Cont) resolveConstructor(v reflect.Value, rt reflect.Type) error {
	if method, found := rt.MethodByName("Construct"); found {
		passParams := []reflect.Value{v}
		methodType := method.Type
		numParams := methodType.NumIn()

		for i := 1; i < numParams; i++ {
			paramType := methodType.In(i)
			param, fullTypeName, err := t.resolveConstructorParam(paramType)
			if err != nil {
				return err
			}
			if t.callStack[fullTypeName] {
				return fmt.Errorf("%w: circular call: %s", ErrCircularReference, fullTypeName)
			}
			t.callStack[fullTypeName] = true

			if _, err := t.getRecursive(param); err != nil {
				delete(t.callStack, fullTypeName)
				return err
			}
			delete(t.callStack, fullTypeName)
			passParams = append(passParams, reflect.ValueOf(param))
		}

		method.Func.Call(passParams)
	}

	return nil
}

// resolveAutoWire resolves dependencies for struct fields with the "di" tag.
func (t *Cont) resolveAutoWire(v reflect.Value) error {
	vTyp := v.Elem().Type()
	for i := 0; i < v.Elem().NumField(); i++ {
		field := vTyp.Field(i)
		tag := field.Tag.Get("di")
		if tag == "autowire" {
			resolvedField := v.Elem().FieldByName(field.Name)

			if !resolvedField.CanSet() {
				return fmt.Errorf("%w: the field name: %s", ErrCannotBeResolvedPossibleNeedExport, field.Name)
			}

			value, _, err := t.resolveConstructorParam(field.Type)
			if err != nil {
				return err
			}

			if _, err := t.getRecursive(value); err != nil {
				return err
			}

			fieldValue := reflect.ValueOf(value)
			if field.Type.Kind() == reflect.Interface && !fieldValue.Type().Implements(field.Type) {
				return fmt.Errorf("provided value does not implement the field's interface: %s", field.Name)
			}

			resolvedField.Set(fieldValue)
		}
	}
	return nil
}

// resolveConstructorParam resolves a constructor parameter by its type.
func (t *Cont) resolveConstructorParam(paramType reflect.Type) (interface{}, string, error) {
	pkgPath := paramType.PkgPath() + "/" + paramType.Name()
	fullTypeName := strings.Join(strings.Split(pkgPath, "/")[1:], ".")
	param, ok := t.dependencies[fullTypeName]
	if !ok {
		return nil, fullTypeName, fmt.Errorf("%w: dependency name: %s", ErrCannotBeResolved, fullTypeName)
	}

	return param, fullTypeName, nil
}
