# Dependency Injection

## Golang Dependency Injection Container

> For developers who miss dependency injection as seen in object-oriented languages.

This container introduces a "Constructor" mechanism for Golang structures. These constructors are automatically called, and their parameters are resolved recursively based on type hints.

Additionally, the container supports struct autowiring.

---

## Usage: Constructor Dependency Resolution

### Create Your Dependency Container:
```go
container := godi.New()
```

### Example Constructor:
```go
func (t *yourStruct) Construct(param yourInterface) {}
```

### Limitations:
- Works only with receiver functions (subject to change in the future).
- Parameters must be interfaces, which will be resolved as either interfaces or structs.

### Mapping Dependencies:
```go
container.Set("TestInterface", NewTest2())
```
- If the interface is in the root folder, use only its name.
- If in a subfolder, provide the full path, e.g., `examplemodule-1.mod.ExampleInterface` (where the interface is in `./examplemodule-1/mod/`).
- If your project has a domain (e.g., `github.com/olbrichattila/godi`), use the module-relative path: `olbrichattila.godi.internal.test.container.noParamConstructorInterface`.

For more examples, check the test folder.

### Resolving Dependencies:
```go
test := NewTest()
_, err := container.Get(test)
if err != nil {
    fmt.Println(err)
}
test.DoWhatYouWant()
```
**Note:** The container returns the struct as `interface{}`; typecast it if needed.

---

## Autowire Usage

To enable autowiring, annotate struct fields:
```go
type exampleMultipleImpl struct {
    Dependency1 exampleMultipleDepInterface `di:"autowire"`
    Dependency2 exampleMultipleDepInterface `di:"autowire"`
    Dependency3 exampleMultipleDepInterface `di:"autowire"`
}
```
Dependencies will be injected automatically when the struct is initialized with the DI container.

---

## Managing Dependencies

### Adding a Single Dependency:
```go
container := godi.New()
container.Set("TestInterface", NewTest2())
```

### Retrieving a Single Dependency:
```go
dependency, err := container.GetDependency("dependencyInterfaceName2")
```

### Adding Multiple Dependencies:
```go
container := godi.New()

dependencies := map[string]interface{}{
    "dependencyInterfaceName1": NewDependency1(),
    "dependencyInterfaceName2": NewDependency2(),
    "dependencyInterfaceName3": NewDependency3(),
}

container.Build(dependencies)
```

### Resolving Dependencies and Calling a Function:
```go
container := godi.New()

dependencies := map[string]interface{}{
    "dependencyInterfaceName1": NewDependency1(),
    "dependencyInterfaceName2": NewDependency2(),
    "dependencyInterfaceName3": NewDependency3(),
}

container.Build(dependencies)

result, err := container.Call(TestFunc)
// result is a slice of reflect.Value

func TestFunc(d1 dependencyInterfaceName1, d2 dependencyInterfaceName2, d3 dependencyInterfaceName3) {}
```

### Resolving Dependencies with Custom Values:
```go
container := godi.New()

dependencies := map[string]interface{}{
    "dependencyInterfaceName1": NewDependency1(),
    "dependencyInterfaceName2": NewDependency2(),
    "dependencyInterfaceName3": NewDependency3(),
}

container.Build(dependencies)

result, err := container.Call(TestFunc, 55, "Hello")
// result is a slice of reflect.Value

func TestFunc(intValue int, stringValue string, d1 dependencyInterfaceName1, d2 dependencyInterfaceName2, d3 dependencyInterfaceName3) {}
```

### Non-Singleton Resolution:
By default, dependencies are treated as singletons. If you need a new instance each time, bind a function that returns a new object:
```go
func yourFunction() {
    container := godi.New()
    oneParamConstructorMock := newOneParamConstClass()

    container.Set("olbrichattila.godi.internal.test.container.noParamConstructorInterface", callerFunc)

    _, err := container.Get(oneParamConstructorMock)
    ...
}

func callerFunc() interface{} {
    return Your()
}
```

### Flushing All Dependencies:
```go
container.Flush()
```

### Deleting a Single Dependency:
```go
container.Delete("dependencyInterfaceName2")
```

### Getting the Number of Dependencies:
```go
count := container.Count()
fmt.Println(count)
```

---

## Example Usage

### `main.go`
```go
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

### `examplemodule-1/mod/example1.go`
```go
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

type example struct {}

func (t *example) Construct(e example2.ExampleInterface) {
    fmt.Println("Constructor of example1 package called")
}
```

### `examplemodule-2/example2.go`
```go
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

type example struct {}

func (t *example) Construct() {
    fmt.Println("Constructor of example2 package called")
}
```

---

## About Me
- Learn more about me on my [personal website](https://attilaolbrich.co.uk/menu/my-story).
- Check out my latest [blog post](https://attilaolbrich.co.uk/blog/1/single).

---

