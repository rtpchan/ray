package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Ray struct {
	Origin *mat.VecDense
	Dir    *mat.VecDense
}

// Define origin x, y, z, and direction i, j k
func NewRay(x, y, z, i, j, k float64) *Ray {
	o := PointV(x, y, z)
	d := NormaliseV(VectorV(i, j, k))
	return &Ray{Origin: o, Dir: d}
}

func (r *Ray) Position(t float64) *mat.VecDense {
	return AddV(r.Origin, ScaleV(t, r.Dir))
}

func IntersectRaySphere(r *Ray, x, y, z float64) (float64, float64) {
	rs := SubV(r.Origin, PointV(x, y, z))
	a := DotV(r.Dir, r.Dir)
	b := 2. * DotV(r.Dir, rs)
	c := DotV(rs, rs) - 1
	discriminant := b*b - 4*a*c
	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	return t1, t2
}
