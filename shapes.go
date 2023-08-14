package main

import "gonum.org/v1/gonum/mat"

type Shape interface {
	ShapeID() int
}

type Sphere struct {
	ID        int
	Origin    *mat.VecDense
	Radius    float64
	Transform *mat.Dense
}

func NewSphere(x, y, z, r float64) *Sphere {
	o := PointV(x, y, z)

	return &Sphere{ID: idgen.NewID(), Origin: o, Radius: r, Transform: Identity()}
}

func (s *Sphere) ShapeID() int {
	return s.ID
}

func (s *Sphere) SetTransform(m *mat.Dense) {
	s.Transform = m
}

func (s *Sphere) NormalAt(p *mat.VecDense) *mat.VecDense {
	objectP := MulV(InverseM(s.Transform), p)
	return NormaliseV(SubV(objectP, s.Origin))
}
