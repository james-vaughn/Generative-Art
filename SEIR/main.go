package main

import (
	"fmt"

	"github.com/james-vaughn/Generative-Art/SEIR/seir"
)

func main() {
	config := seir.SeirConfig{
		Beta:  .8,
		Gamma: .1,
		Sigma: .4,
		Mu:    0,
		Nu:    0,
		S:     600,
		E:     1,
		I:     2,
		R:     0,
		Days:  40,
		Steps: 100,
	}

	model := seir.NewModel(config)
	model.CalculatePoints()

	fmt.Println(model.S)
	fmt.Println(model.E)
	fmt.Println(model.I)
	fmt.Println(model.R)
}
