package test

func newOneParamConstructorMock() oneParamRecursionConstructorInterface {
	return &oneParamConstructor{}
}

type oneParamRecursionConstructorInterface interface {
	Construct(noParamConstructorInterface)
	ConstructorCallCount() int
}

type oneParamConstructor struct {
	constructorCalled int
}

func (t *oneParamConstructor) Construct(_ noParamConstructorInterface) {
	t.constructorCalled++
}

func (t *oneParamConstructor) ConstructorCallCount() int {
	return t.constructorCalled
}
