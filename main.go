package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

var (
	idgen *ID
)

func init() {
	idgen = NewIDGenerator()
}

func main() {
	fmt.Println("ray tracing")

	a := mat.NewVecDense(3, []float64{1, 2, 3})
	b := mat.NewVecDense(3, []float64{4, 5, 6})
	var c mat.VecDense
	c.AddVec(a, b)
	fmt.Println(c)
}
