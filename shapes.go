package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Shaper interface {
	ShapeID() int
	NormalAt(*mat.VecDense) *mat.VecDense // copy same function to each shape
	GetTransform() *mat.Dense
	GetMaterial() *Material
	Intersect(*Ray) Intersections // run LocalRay in each shape
	LocalNormalAt(*mat.VecDense) *mat.VecDense
}

type Shape struct {
	ID        int
	Transform *mat.Dense
	Material  *Material
	SavedRay  *Ray
}

func NewShape() Shape {
	return Shape{ID: idgen.NewID(), Transform: Identity(), Material: NewMaterial()}
}

func (s *Shape) ShapeID() int {
	return s.ID
}

func (s *Shape) GetTransform() *mat.Dense {
	return s.Transform
}

func (s *Shape) SetTransform(m *mat.Dense) {
	s.Transform = m
}

func (s *Shape) GetMaterial() *Material {
	return s.Material
}

func (s *Shape) SetMaterial(m *Material) {
	s.Material = m
}

func (s *Shape) LocalRay(r *Ray) *Ray {
	invM := ZeroMatrix()
	invM.Inverse(s.GetTransform())
	s.SavedRay = r.Transform(invM)
	return s.SavedRay
}

type Sphere struct {
	Origin *mat.VecDense
	Radius float64
	Shape
}

func NewSphere() *Sphere {
	o := PointV(0, 0, 0)
	return &Sphere{Shape: NewShape(), Origin: o, Radius: 1.0}
}

func (s *Sphere) GetOrigin() *mat.VecDense {
	return s.Origin
}

// Get normal of a point on a shape, see Chapter 6 Page 82
func (s *Sphere) NormalAt(p *mat.VecDense) *mat.VecDense {
	objectP := MulV(InverseM(s.Transform), p)
	objectN := s.LocalNormalAt(objectP)
	worldN := MulTranposeMV(InverseM(s.Transform), objectN)
	worldN.SetVec(3, 0.0)
	return NormaliseV(worldN)
}

func (s *Sphere) LocalNormalAt(lp *mat.VecDense) *mat.VecDense {
	return SubV(lp, s.Origin)
}

func (s *Sphere) Intersect(r *Ray) Intersections {
	r = s.LocalRay(r)
	rs := SubV(r.Origin, s.GetOrigin())
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

type Plane struct {
	Shape
}

func NewPlane() *Plane {

	return &Plane{Shape: NewShape()}
}

func (s *Plane) Intersect(r *Ray) Intersections {
	r = s.LocalRay(r)
	if math.Abs(r.Dir.AtVec(1)) < EPSILON {
		return []Intersection{}
	}
	t := -r.Origin.AtVec(1) / r.Dir.AtVec(1)
	return []Intersection{{t, s}}
}

// Get normal of a point on a shape, see Chapter 6 Page 82
func (s *Plane) NormalAt(p *mat.VecDense) *mat.VecDense {
	objectP := MulV(InverseM(s.Transform), p)
	objectN := s.LocalNormalAt(objectP)
	worldN := MulTranposeMV(InverseM(s.Transform), objectN)
	worldN.SetVec(3, 0.0)
	return NormaliseV(worldN)
}

func (s *Plane) LocalNormalAt(p *mat.VecDense) *mat.VecDense {
	return VectorV(0, 1, 0)
}
