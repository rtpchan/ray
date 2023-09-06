package main

import (
	"fmt"
	"log"
	"time"
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
	start := time.Now()
	Chapter10()
	duration := time.Since(start)
	log.Printf("Duration %0.2fs\n", duration.Seconds())
}
