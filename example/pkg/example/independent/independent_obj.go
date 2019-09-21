package independent

type Obj struct {
	SomeValue string
}

func NewIndependentObj(someValue string) *Obj {
	return &Obj{someValue}
}
