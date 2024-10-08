# Dependency injection

## Golang dependency injection container

> For those who are missing a Dependency injection similar to what we use in some object oriented languages.

> This container introduces a "Constructor" for golang structures, Those constructors will automatically be called and the 
parameters will be resolved by their type hint recursively.

> This container also introduces struct autowire notation.

## Usage: Constructor dependency resolutiuon

### Create your dependency container:
```
container := godi.New()
```

### Example constructor:
```
func (t *yourStruct) Construct(param yourInterfaceInterface) {}
```

### Limitation:
> It works only with receiver functions, this may change in the future.

The parameter type always have to be an interface, the parameter is what have to be resolved, That can be an interface or a struct.

### Map your dependencies:
```
container.Set("TestInterface", NewTest2())
```

Where "TestInterface" is your interface name, "NewTest2()" returns an interface or a struct. If a Constructor (exported) method is defined it will be called and the hinted dependencies will be resolved from the map (please see later)

- If it is in defined in your home folder, then use only the interface name
- If it is in a sub folder, provide full path, exampe: ```examplemodule-1.mod.ExampleInterface``` where your interface defined in folder ```./examplemodule-1/mod/```
- If you initiated your project with a domain. (for example this module) ```github.com/olbrichattila/godi```. use the path from your module as well. Example: ```olbrichattila.godi.internal.test.container.noParamConstructorInterface```

Note: Look at the test folder for examples, and see the following example as well:

### Resolve your dependencies from your container instance:
```
test := NewTest
_, err := container.Get(test)
	if err != nil {
		fmt.Println(err)
	}
test.DoWhatYouWant()
```

Note: The container also returns the original (test) struct, but as interface{}. if you use this, please typecast it back. (_) parameter

### Usage autowire 
Add annotation to fields like: 
```
type exampleMultipleImpl struct {
	Dependency1 exampleMultipleDepInterface `di:"autowire"`
	Dependency2 exampleMultipleDepInterface `di:"autowire"`
	Dependency3 exampleMultipleDepInterface `di:"autowire"`
}
```

The dependencies will be auto wired, if you initiate your struct with the DI container as above:

## Managing dependencies
One way of managing dependencies was already introduced above:

### Add single dependency
```
container := godi.New()

container.Set("TestInterface", NewTest2())
```
 ### Retrieve single dependency

```
dependency, err := container.GetDependency("dependencyInterfaceName2")
```

### Add multiple dependencies at the same time
```
container := godi.New()

dependencies := map[string]interface{}{
	"dependencyInterfaceName1": NewDependency1(),
	"dependencyInterfaceName2": NewDependency1(),
	"dependencyInterfaceName3": NewDependency1(),
}

container.Build(dependencies)
```

### Resolve dependencies and call a single function
```
container := godi.New()

dependencies := map[string]interface{}{
	"dependencyInterfaceName1": NewDependency1(),
	"dependencyInterfaceName2": NewDependency1(),
	"dependencyInterfaceName3": NewDependency1(),
}

container.Build(dependencies)

result, err := container.Call(TestFunc)
// Where result is a slice of Reflect.Value

func TestFunc(d1 dependencyInterfaceName1, d2 dependencyInterfaceName2, d3 dependencyInterfaceName3) {
}
```

### Resolve dependencies and add custom values and call a single function
```
container := godi.New()

dependencies := map[string]interface{}{
	"dependencyInterfaceName1": NewDependency1(),
	"dependencyInterfaceName2": NewDependency1(),
	"dependencyInterfaceName3": NewDependency1(),
}

container.Build(dependencies)

result, err := container.Call(TestFunc, 55, "Hello")
// Where result is a slice of Reflect.Value

func TestFunc(intValue int, stringValue string, d1 dependencyInterfaceName1, d2 dependencyInterfaceName2, d3 dependencyInterfaceName3) {
}
```

### Non singleton resolution
The above examples are always returning the initiated struct (like a singleton)
If you would like to have a new instance, use a function to create the binding with the function.
The function must follow ```func() interface{}``` always
Example:
```
func yourFunction() {
	container := godi.New()
	oneParamConstructorMock := newOneParamConstClass()

	container.Set("olbrichattila.godi.internal.test.container.noParamConstructorInterface", callerFunc)

	_, err := container.Get(oneParamConstructorMock)
	...
}

func callerFunc() interface{} {
	return YOur()
}
```

### Flush all dependencies
```
container.Flush()
```

### Delete one dependency
```
container.Delete("dependencyInterfaceName2")
```

### Get number or existing dependencies
```
count := container.Count()
fmt.Println(count)
```


Example usage:
```main.go```

```
package main

import (
	"fmt"
	example1 "godi-test/examplemodule-1/mod"
	example2 "godi-test/examplemodule-2"

	"github.com/olbrichattila/godi"
)

func main() {
	container := godi.New()
	container.Set("TestInterface2", NewTest2())
	container.Set("examplemodule-1.mod.ExampleInterface", example1.New())
	container.Set("examplemodule-2.ExampleInterface", example2.New())

	_, err := container.Get(NewTest())
	if err != nil {
		fmt.Println(err)
	}
}

func NewTest() TestInterface {
	return &test{}
}

type TestInterface interface {
	Construct(TestInterface2, example1.ExampleInterface)
}

type test struct{}

func (*test) Construct(i TestInterface2, ex1 example1.ExampleInterface) {
	fmt.Println("Hello from construct")
}

func NewTest2() TestInterface2 {
	return &test2{}
}

type TestInterface2 interface {
	Construct()
}

type test2 struct{}

func (*test2) Construct() {
	fmt.Println("Hello from construct 2")
}

```
```examplemodule-1/mod/example1.go```

```
package example1

import (
	"fmt"
	example2 "godi-test/examplemodule-2"
)

func New() ExampleInterface {
	return &example{}
}

type ExampleInterface interface {
	Construct(example2.ExampleInterface)
}

type example struct {
}

func (t *example) Construct(e example2.ExampleInterface) {
	fmt.Println("Constructor of example1 package called")
}

```
```examplemodule-2/example2.go```
```
package example2

import (
	"fmt"
)

func New() ExampleInterface {
	return &example{}
}

type ExampleInterface interface {
	Construct()
}

type example struct {
}

func (t *example) Construct() {
	fmt.Println("Constructor of example2 package called")
}
```

