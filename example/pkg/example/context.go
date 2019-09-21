package example

import (
	"github.com/catmorte/go-inversion_of_control/pkg/context"
)

func init() {
	context.SetContext(context.NewMemoryContext())
}
