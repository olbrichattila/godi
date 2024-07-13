package test

func newNestedSecondMock() nestedSecondInterface {
	return &nestedSecond{}
}

type nestedSecondInterface interface {
	Construct(nestedThirdInterface)
	ConstructorCallCount() int
	MockFunc()
	MockFuncCallCount() int
}

type nestedSecond struct {
	nestedInstance    nestedThirdInterface
	constructorCalled int
	mockFuncCalled    int
}

func (t *nestedSecond) Construct(n nestedThirdInterface) {
	t.nestedInstance = n
	t.constructorCalled++
}

func (t *nestedSecond) ConstructorCallCount() int {
	return t.constructorCalled
}

func (t *nestedSecond) MockFunc() {
	t.nestedInstance.MockFunc()
	t.mockFuncCalled++
}

func (t *nestedSecond) MockFuncCallCount() int {
	return t.mockFuncCalled
}
