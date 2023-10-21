package main

import (
	"fmt"

	_ "github.com/catmorte/go-inversion_of_control/example/pkg/example/beans"
	"github.com/catmorte/go-inversion_of_control/example/pkg/example/dependent"
	. "github.com/catmorte/go-inversion_of_control/pkg/context"
)

func main() {
	dependentBean := Ask[*dependent.Obj]()
	fmt.Println(dependentBean.IndependentObj.SomeValue)
}
