package main

import (
	"fmt"
)

const (
	EPSILON float64 = 0.00001
)

var (
	idgen *ID
)

func init() {
	idgen = NewIDGenerator()
}

func main() {
	fmt.Println("ray tracing")
	Chapter9()
}
