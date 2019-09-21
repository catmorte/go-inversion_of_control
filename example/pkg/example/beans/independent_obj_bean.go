package beans

import (
	"github.com/catmorte/go-inversion_of_control/example/pkg/example/independent"
	"github.com/catmorte/go-inversion_of_control/pkg/context"
)

func init() {
	context.GetContext().Reg((*independent.Obj)(nil), func() interface{} {
		return independent.NewIndependentObj("Hello world!")
	})
}
