package test

func newNestedCircularFirstMock() nestedCircularFirstInterface {
	return &nestedCircularFirst{}
}

type nestedCircularFirstInterface interface {
	Construct(nestedCircularSecondInterface)
}

type nestedCircularFirst struct {
}

func (t *nestedCircularFirst) Construct(_ nestedCircularSecondInterface) {
}
