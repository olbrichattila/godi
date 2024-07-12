# Dependency injection

Work in progress, usage description to come

Example:
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

TODO:
- Check circular dependencies and throw error
- Test
- Struct autowire

