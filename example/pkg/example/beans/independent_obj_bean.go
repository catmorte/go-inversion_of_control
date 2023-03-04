package beans

import (
	"github.com/catmorte/go-inversion_of_control/example/pkg/example/independent"
	"github.com/catmorte/go-inversion_of_control/pkg/context"
)

func init() {
	context.Reg[*independent.Obj](context.GetContext(), func() interface{} {
		return independent.NewIndependentObj("Hello world!")
	})
}
