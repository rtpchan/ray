package main

import "gonum.org/v1/gonum/mat"

func PointM(x, y, z float64) *mat.VecDense {
	return mat.NewVecDense(4, []float64{x, y, z, 1})
}

func VectorM(x, y, z float64) *mat.VecDense {
	return mat.NewVecDense(4, []float64{x, y, z, 0})
}

func IsPointM(v *mat.VecDense) bool {
	return v.AtVec(3) == 1
}

func IsVectorM(v *mat.VecDense) bool {
	return v.AtVec(3) == 0
}

func MagnitudeM(v *mat.VecDense) float64 {
	return v.Norm(2)
}

func NormalisationM(v *mat.VecDense) *mat.VecDense {
	// c := VectorM(0, 0, 0)
	var c mat.VecDense
	c.ScaleVec(1./v.Norm(2), v)
	return &c
}

func DotM(a, b *mat.VecDense) float64 {
	return mat.Dot(a, b)
}

func CrossM(a, b *mat.VecDense) *mat.VecDense {
	c := VectorM(0, 0, 0)
	c.SetVec(0, a.AtVec(1)*b.AtVec(2)-a.AtVec(2)*b.AtVec(1))
	c.SetVec(1, a.AtVec(2)*b.AtVec(0)-a.AtVec(0)*b.AtVec(2))
	c.SetVec(2, a.AtVec(0)*b.AtVec(1)-a.AtVec(1)*b.AtVec(0))

	return c
}
