package main

import "gonum.org/v1/gonum/mat"

func PointV(x, y, z float64) *mat.VecDense {
	return mat.NewVecDense(4, []float64{x, y, z, 1})
}

func VectorV(x, y, z float64) *mat.VecDense {
	return mat.NewVecDense(4, []float64{x, y, z, 0})
}

func IsPointV(v *mat.VecDense) bool {
	return v.AtVec(3) == 1
}

func IsVectorV(v *mat.VecDense) bool {
	return v.AtVec(3) == 0
}

func AddV(a, b *mat.VecDense) *mat.VecDense {
	var c mat.VecDense
	c.AddVec(a, b)
	return &c
}

func SubV(a, b *mat.VecDense) *mat.VecDense {
	var c mat.VecDense
	c.SubVec(a, b)
	return &c
}

func ScaleV(a float64, b *mat.VecDense) *mat.VecDense {
	var c mat.VecDense
	c.ScaleVec(a, b)
	return &c
}

func MagnitudeV(v *mat.VecDense) float64 {
	return v.Norm(2)
}

func NormalisationV(v *mat.VecDense) *mat.VecDense {
	// c := VectorM(0, 0, 0)
	var c mat.VecDense
	c.ScaleVec(1./v.Norm(2), v)
	return &c
}

func DotV(a, b *mat.VecDense) float64 {
	return mat.Dot(a, b)
}

func CrossV(a, b *mat.VecDense) *mat.VecDense {
	c := VectorV(0, 0, 0)
	c.SetVec(0, a.AtVec(1)*b.AtVec(2)-a.AtVec(2)*b.AtVec(1))
	c.SetVec(1, a.AtVec(2)*b.AtVec(0)-a.AtVec(0)*b.AtVec(2))
	c.SetVec(2, a.AtVec(0)*b.AtVec(1)-a.AtVec(1)*b.AtVec(0))

	return c
}
