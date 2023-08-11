package main

import (
	"math"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestRay(t *testing.T) {
	r := NewRay(2, 3, 4, 1, 0, 0)
	a := r.Position(1)
	ar := PointV(3, 3, 4)
	if !mat.EqualApprox(a, ar, 0.000001) {
		t.Errorf("Ray position should be %v, got %v", ar, a)
	}

	b := r.Position(-1)
	br := PointV(1, 3, 4)
	if !mat.EqualApprox(b, br, 0.000001) {
		t.Errorf("Ray position should be %v, got %v", br, b)
	}
}

func TestRaySphere(t *testing.T) {
	r := NewRay(0, 0, 0, 0, 0, 1)
	x1, x2 := IntersectRaySphere(r, 0, 0, 0)
	if x1 != -1 || x2 != 1 {
		t.Errorf("Ray Sphere intersection at -1, 1, got %0.2f, %0.2f", x1, x2)
	}
	r2 := NewRay(0, 0, 5, 0, 0, 1)
	x3, x4 := IntersectRaySphere(r2, 0, 0, 0)
	if x3 != -6. || x4 != -4.0 {
		t.Errorf("Ray Sphere intersection at -6, -4, got %0.2f, %0.2f", x3, x4)
	}
	r3 := NewRay(0, 2, -5, 0, 0, 1)
	x5, x6 := IntersectRaySphere(r3, 0, 0, 0)
	if !math.IsNaN(x5) || !math.IsNaN(x6) {
		t.Errorf("Ray Sphere intersection at -6, -4, got %0.2f, %0.2f", x5, x6)
	}
}
