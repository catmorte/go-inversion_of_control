package dependent

import (
	"github.com/catmorte/go-inversion_of_control/example/pkg/example/independent"
)

type Obj struct {
	IndependentObj *independent.Obj
}

func NewDependentObj(independentDep *independent.Obj) *Obj {
	return &Obj{independentDep}
}
