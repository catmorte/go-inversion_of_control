package context

import (
	"testing"
)

func init() {
	SetContext(NewMemoryContext())
}

type firstIndependentStruct struct {
	val string
}
type secondIndependentStruct struct {
	val string
}
type dependentStruct struct {
	firstDep  *firstIndependentStruct
	secondDep *secondIndependentStruct
}

func TestMemoryContext(t *testing.T) {
	GetContext().Reg((*firstIndependentStruct)(nil), func() interface{} {
		t.Log("Start init firstIndependentStruct")
		return &firstIndependentStruct{"firstTestString"}
	})

	firstDep := Dep((*firstIndependentStruct)(nil))
	secondDep := Dep((*secondIndependentStruct)(nil))
	GetContext().Reg((*dependentStruct)(nil), func() interface{} {
		t.Log("Start init dependentStruct")
		return &dependentStruct{
			firstDep:  (<-firstDep.Waiter).(*firstIndependentStruct),
			secondDep: (<-secondDep.Waiter).(*secondIndependentStruct),
		}
	}, firstDep, secondDep)

	GetContext().Reg((*secondIndependentStruct)(nil), func() interface{} {
		t.Log("Start init secondIndependentStruct")
		return &secondIndependentStruct{"secondTestString"}
	})

	t.Log("Start waiting for dependentStruct")
	actualInst := (<-GetContext().Ask((*dependentStruct)(nil))).(*dependentStruct)
	if actualInst.firstDep.val == "firstTestString" && actualInst.secondDep.val == "secondTestString" {
		t.Log("Initialized")
		return
	}
	t.Errorf("Expected values %v %v", "firstTestString", "secondTestString")
}
