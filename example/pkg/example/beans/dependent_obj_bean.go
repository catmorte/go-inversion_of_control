package beans

import (
	"github.com/catmorte/go-inversion_of_control/example/pkg/example/dependent"
	"github.com/catmorte/go-inversion_of_control/example/pkg/example/independent"
	"github.com/catmorte/go-inversion_of_control/pkg/context"
)

func init() {
	independentDep := context.Dep((*independent.Obj)(nil))
	context.GetContext().Reg((*dependent.Obj)(nil), func() interface{} {
		return dependent.NewDependentObj((<-independentDep.Waiter).(*independent.Obj))
	}, independentDep)
}
