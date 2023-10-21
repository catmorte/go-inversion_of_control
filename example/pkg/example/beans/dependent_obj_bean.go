package beans

import (
	"github.com/catmorte/go-inversion_of_control/example/pkg/example/dependent"
	"github.com/catmorte/go-inversion_of_control/example/pkg/example/independent"
	. "github.com/catmorte/go-inversion_of_control/pkg/context"
)

func init() {
	independentDep := Dep[*independent.Obj]()
	Reg[*dependent.Obj](func() interface{} {
		return dependent.NewDependentObj(ResolveDep[*independent.Obj](independentDep))
	}, independentDep)
}
