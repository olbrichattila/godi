package containertest

func newTreeParamConstructorMock() threeParamRecursionConstructorInterface {
	return &threeParamConstructor{}
}

type threeParamRecursionConstructorInterface interface {
	Construct(noParamConstructorInterface, noParamConstructorInterface, noParamConstructorInterface)
	ConstructorCallCount() int
}

type threeParamConstructor struct {
	constructorCalled int
}

func (t *threeParamConstructor) Construct(_ noParamConstructorInterface, _ noParamConstructorInterface, _ noParamConstructorInterface) {
	t.constructorCalled++
}

func (t *threeParamConstructor) ConstructorCallCount() int {
	return t.constructorCalled
}
