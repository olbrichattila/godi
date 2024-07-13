package test

func newNestedThirdMock() nestedThirdInterface {
	return &nestedThird{}
}

type nestedThirdInterface interface {
	Construct()
	ConstructorCallCount() int
	MockFunc()
	MockFuncCallCount() int
}

type nestedThird struct {
	constructorCalled int
	mockFuncCalled    int
}

func (t *nestedThird) Construct() {
	t.constructorCalled++
}

func (t *nestedThird) ConstructorCallCount() int {
	return t.constructorCalled
}

func (t *nestedThird) MockFunc() {
	t.mockFuncCalled++
}

func (t *nestedThird) MockFuncCallCount() int {
	return t.mockFuncCalled
}
