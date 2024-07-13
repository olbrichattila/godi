package test

func newNestedCircularSecondMock() nestedCircularSecondInterface {
	return &nestedCircularSecond{}
}

type nestedCircularSecondInterface interface {
	Construct(nestedCircularThirdInterface)
}

type nestedCircularSecond struct {
}

func (t *nestedCircularSecond) Construct(_ nestedCircularThirdInterface) {
}
