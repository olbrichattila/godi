package test

func newNestedCircularMock() nestedCircularInterface {
	return &nestedCircular{}
}

type nestedCircularInterface interface {
	Construct(nestedCircularInterface)
}

type nestedCircular struct {
}

func (t *nestedCircular) Construct(n nestedCircularInterface) {
}
