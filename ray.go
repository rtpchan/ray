package main

import (
	"gonum.org/v1/gonum/mat"
)

type Ray struct {
	Origin *mat.VecDense
	Dir    *mat.VecDense
}

// Define origin x, y, z, and direction i, j k
func NewRayCoor(x, y, z, i, j, k float64) *Ray {
	o := PointV(x, y, z)
	d := NormaliseV(VectorV(i, j, k)) // do not normalise, (for world space ?, page 69)
	// return &Ray{Origin: o, Dir: VectorV(i, j, k)}
	return &Ray{Origin: o, Dir: d}

}

func NewRay(o *mat.VecDense, d *mat.VecDense) *Ray {
	return &Ray{Origin: o, Dir: d}
}

func (r *Ray) Position(t float64) *mat.VecDense {
	return AddV(r.Origin, ScaleV(t, r.Dir))
}

func (r *Ray) Transform(m *mat.Dense) *Ray {
	np := PointV(0, 0, 0)
	np.MulVec(m, r.Origin)
	nv := VectorV(0, 0, 0)
	nv.MulVec(m, r.Dir)
	return NewRay(np, nv)
}
