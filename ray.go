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
func NewRayCoor(x, y, z, i, j, k float64) *Ray {
	o := PointV(x, y, z)
	// d := NormaliseV(VectorV(i, j, k))   // do not normalise, (for world space ?, page 69)
	return &Ray{Origin: o, Dir: VectorV(i, j, k)}
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

func IntersectRaySphere(r *Ray, s *Sphere) []Intersection {
	invM := ZeroMatrix()
	invM.Inverse(s.Transform)
	r = r.Transform(invM)

	rs := SubV(r.Origin, s.Origin)
	a := DotV(r.Dir, r.Dir)
	b := 2. * DotV(r.Dir, rs)
	c := DotV(rs, rs) - 1
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return []Intersection{}
	}
	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	return []Intersection{
		{T: t1, Object: s},
		{T: t2, Object: s},
	}
}
