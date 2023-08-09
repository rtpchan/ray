package main

// Vector function on []float64, assume 3 dimensions, i.e. up to index 2
// Assumes 3 dimensions,

import (
	"math"
)

func Vector(x, y, z float64) [3]float64 {
	return [3]float64{x, y, z}
}

func Add(a, b [3]float64) [3]float64 {
	return [3]float64{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

func Sub(a, b [3]float64) [3]float64 {
	return [3]float64{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func Scale(a [3]float64, s float64) [3]float64 {
	return [3]float64{a[0] * s, a[1] * s, a[2] * s}

}

func Negate(a [3]float64) [3]float64 {
	return Scale(a, -1)
}

func Magnitude(a [3]float64) float64 {
	return math.Sqrt(a[0]*a[0] + a[1]*a[1] + a[2]*a[2])
}

func Normalise(a [3]float64) [3]float64 {
	m := Magnitude(a)
	return [3]float64{a[0] / m, a[1] / m, a[2] / m}
}

func Dot(a, b [3]float64) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

func Cross(a, b [3]float64) [3]float64 {
	return [3]float64{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
	}
}

func EqualVector(a, b [3]float64, tol float64) bool {
	for i := range a {
		if math.Abs(a[i]-b[i]) > tol {
			return false
		}
	}
	return true
}

func EqualFloat(a, b, tol float64) bool {
	return math.Abs(a-b) < tol
}
