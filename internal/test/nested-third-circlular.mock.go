package test

func newNestedCircularThirdMock() nestedCircularThirdInterface {
	return &nestedCircularThird{}
}

type nestedCircularThirdInterface interface {
	Construct(nestedCircularFirstInterface)
}

type nestedCircularThird struct {
}

func (t *nestedCircularThird) Construct(_ nestedCircularFirstInterface) {
}
