package main

import (
	"fmt"
	_ "github.com/catmorte/go-inversion_of_control/example/pkg/example"
	_ "github.com/catmorte/go-inversion_of_control/example/pkg/example/beans"
	"github.com/catmorte/go-inversion_of_control/example/pkg/example/dependent"
	"github.com/catmorte/go-inversion_of_control/pkg/context"
)

func main() {
	dependentBean := context.Ask[*dependent.Obj](context.GetContext())
	fmt.Println(dependentBean.IndependentObj.SomeValue)
}
