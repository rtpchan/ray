package main

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestRay(t *testing.T) {
	r := NewRayCoor(2, 3, 4, 1, 0, 0)
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
	r := NewRayCoor(0, 0, 0, 0, 0, 1)
	s := NewSphere(0, 0, 0, 2)
	x := IntersectRaySphere(r, s)
	if x[0].T != -1 || x[1].T != 1 {
		t.Errorf("Ray Sphere intersection at -1, 1, got %0.2f, %0.2f", x[0].T, x[1].T)
	}
	r2 := NewRayCoor(0, 0, 5, 0, 0, 1)
	s2 := NewSphere(0, 0, 0, 2)
	x2 := IntersectRaySphere(r2, s2)
	if x2[0].T != -6. || x2[1].T != -4.0 {
		t.Errorf("Ray Sphere intersection at -6, -4, got %0.2f, %0.2f", x2[0].T, x2[1].T)
	}
	r3 := NewRayCoor(0, 2, -5, 0, 0, 1)
	s3 := NewSphere(0, 0, 0, 5)
	x3 := IntersectRaySphere(r3, s3)
	if len(x3) != 0 {
		t.Errorf("Ray Sphere intersection has no intersection, got %v", x3)
	}
}

func TestRayTransform(t *testing.T) {
	r := NewRayCoor(1, 2, 3, 0, 1, 0)
	m := TranslateM(3, 4, 5)
	r2 := r.Transform(m)
	if !mat.EqualApprox(r2.Origin, PointV(4, 6, 8), 0.00001) {
		t.Errorf("Ray Translate, should get %v, got %v", PointV(4, 6, 8), r2.Origin)
	}
	if !mat.EqualApprox(r2.Dir, VectorV(0, 1, 0), 0.000001) {
		t.Errorf("Ray Translate, should get %v, got %v", VectorV(0, 1, 0), r2.Dir)
	}

	m = ScaleM(2, 3, 4)
	r2 = r.Transform(m)
	if !mat.EqualApprox(r2.Origin, PointV(2, 6, 12), 0.00001) {
		t.Errorf("Ray Translate, should get %v, got %v", PointV(2, 6, 12), r2.Origin)
	}
	if !mat.EqualApprox(r2.Dir, VectorV(0, 3, 0), 0.000001) {
		t.Errorf("Ray Translate, should get %v, got %v", VectorV(0, 3, 0), r2.Dir)
	}
}
