package main

import (
	"fmt"
)

var (
	idgen *ID
)

func init() {
	idgen = NewIDGenerator()
}

func main() {
	fmt.Println("ray tracing")
	Chapter5()
}
