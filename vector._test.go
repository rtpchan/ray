package main

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestVector(t *testing.T) {
	a := VectorM(0, 3, 4)
	// m := a.Norm(2)
	m := MagnitudeM(a)
	if m != 5 {
		t.Errorf("Vector Magnitude, should be 5, got %0.2f", m)
	}

	b := VectorM(0, 3./5., 4./5.)
	// c := VectorM(0, 0, 0)
	// c.ScaleVec(1./a.Norm(2), a)
	c := NormalisationM(a)
	if !mat.EqualApprox(c, b, 0.0001) {
		t.Errorf("Vector Normalisation, got %v", c)
	}

	e := VectorM(1, 2, 3)
	f := VectorM(2, 3, 4)
	if mat.Dot(e, f) != 20 {
		t.Errorf("Vector Dot Product, should be 20, got %v", mat.Dot(e, f))
	}

	g := CrossM(e, f)
	h := VectorM(-1, 2, -1)
	if !mat.Equal(g, h) {
		t.Errorf("Vector Cross Product, got %v", g)
	}
	n := CrossM(f, e)
	k := VectorM(1, -2, 1)
	if !mat.Equal(k, n) {
		t.Errorf("Vector Cross Product, got %v", n)
	}
}
