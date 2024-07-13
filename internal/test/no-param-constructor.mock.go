package test

func newNoParamConstructorMock() noParamConstructorInterface {
	return &noParamConstructor{}
}

type noParamConstructorInterface interface {
	Construct()
	ConstructorCallCount() int
}

type noParamConstructor struct {
	constructorCalled int
}

func (t *noParamConstructor) Construct() {
	t.constructorCalled++
}

func (t *noParamConstructor) ConstructorCallCount() int {
	return t.constructorCalled
}
