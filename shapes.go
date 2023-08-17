package main

import "gonum.org/v1/gonum/mat"

type Shape interface {
	ShapeID() int
	NormalAt(*mat.VecDense) *mat.VecDense
	GetOrigin() *mat.VecDense
	GetTransform() *mat.Dense
	GetMaterial() *Material
}

type Sphere struct {
	ID        int
	Origin    *mat.VecDense
	Radius    float64
	Transform *mat.Dense
	Material  *Material
}

func NewSphere(x, y, z, r float64) *Sphere {
	// o := PointV(x, y, z)
	o := PointV(0, 0, 0)

	// return &Sphere{ID: idgen.NewID(), Origin: o, Radius: r, Transform: Identity()}
	return &Sphere{ID: idgen.NewID(), Origin: o, Radius: 1.0, Transform: Identity()}
}

func (s *Sphere) ShapeID() int {
	return s.ID
}

func (s *Sphere) GetOrigin() *mat.VecDense {
	return s.Origin
}

func (s *Sphere) GetTransform() *mat.Dense {
	return s.Transform
}

func (s *Sphere) GetMaterial() *Material {
	return s.Material
}

func (s *Sphere) SetTransform(m *mat.Dense) {
	s.Transform = m
}

// Get normal of a point on a sphere, see Chapter 6 Page 82
func (s *Sphere) NormalAt(p *mat.VecDense) *mat.VecDense {
	objectP := MulV(InverseM(s.Transform), p)
	objectN := SubV(objectP, s.Origin)
	worldN := MulTranposeMV(InverseM(s.Transform), objectN)
	worldN.SetVec(3, 0.0)
	return NormaliseV(worldN)
}
