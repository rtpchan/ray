package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func TranslateM(x, y, z float64) *mat.Dense {
	return mat.NewDense(4, 4, []float64{1, 0, 0, x, 0, 1, 0, y, 0, 0, 1, z, 0, 0, 0, 1})
}

func ScaleM(x, y, z float64) *mat.Dense {
	return mat.NewDense(4, 4, []float64{x, 0, 0, 0, 0, y, 0, 0, 0, 0, z, 0, 0, 0, 0, 1})
}

func RotateXM(r float64) *mat.Dense {
	return mat.NewDense(4, 4, []float64{
		1, 0, 0, 0,
		0, math.Cos(r), -math.Sin(r), 0,
		0, math.Sin(r), math.Cos(r), 0,
		0, 0, 0, 1})
}

func RotateYM(r float64) *mat.Dense {
	return mat.NewDense(4, 4, []float64{
		math.Cos(r), 0, math.Sin(r), 0,
		0, 1, 0, 0,
		-math.Sin(r), 0, math.Cos(r), 0,
		0, 0, 0, 1})
}

func RotateZM(r float64) *mat.Dense {
	return mat.NewDense(4, 4, []float64{
		math.Cos(r), -math.Sin(r), 0, 0,
		math.Sin(r), math.Cos(r), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1})
}

func ShearM(xy, xz, yx, yz, zx, zy float64) *mat.Dense {
	return mat.NewDense(4, 4, []float64{
		1, xy, xz, 0,
		yx, 1, yz, 0,
		zx, zy, 1, 0,
		0, 0, 0, 1})
}

// Apply Transformation matrix to a column vector
func AT(m *mat.Dense, v *mat.VecDense) *mat.VecDense {
	r := ZeroVector()
	r.MulVec(m, v)
	return r
}

// Chain Transformation matrix
func CT(m, n *mat.Dense) *mat.Dense {
	q := ZeroMatrix()
	q.Mul(m, n)
	return q
}
