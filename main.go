package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)


func main() {
	a := mat.NewVecDense(3, []float64{1,2,3})
	b := mat.NewVecDense(3, []float64{4,5,6})
	var c mat.VecDense
	fmt.Println("vim-go")
	c.AddVec(a,b)
	fmt.Println(c)
}

func add(a, b int) int {
	return a + b
}
