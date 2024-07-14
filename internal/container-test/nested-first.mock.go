package containertest

func newNestedFirstMock() nestedFirstInterface {
	return &nestedFirst{}
}

type nestedFirstInterface interface {
	Construct(nestedSecondInterface)
	ConstructorCallCount() int
	MockFunc()
	MockFuncCallCount() int
}

type nestedFirst struct {
	nestedInstance    nestedSecondInterface
	constructorCalled int
	mockFuncCalled    int
}

func (t *nestedFirst) Construct(n nestedSecondInterface) {
	t.nestedInstance = n
	t.constructorCalled++
}

func (t *nestedFirst) ConstructorCallCount() int {
	return t.constructorCalled
}

func (t *nestedFirst) MockFunc() {
	t.nestedInstance.MockFunc()
	t.mockFuncCalled++
}

func (t *nestedFirst) MockFuncCallCount() int {
	return t.mockFuncCalled
}
