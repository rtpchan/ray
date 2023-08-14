package main

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestSphereTransform(t *testing.T) {
	r := NewRayCoor(0, 0, -5, 0, 0, 1)
	s := NewSphere(0, 0, 0, 1)
	m := ScaleM(2, 2, 2)
	s.SetTransform(m)
	xs := IntersectRaySphere(r, s)
	if xs[0].T != 3 {
		t.Errorf("Sphere Transform t should be 3, got %0.2f", xs[0].T)
	}
	if xs[1].T != 7 {
		t.Errorf("Sphere Transform t should be 7, got %0.2f", xs[1].T)
	}

	s = NewSphere(0, 0, 0, 1)
	m = TranslateM(5, 0, 0)
	s.SetTransform(m)
	xs = IntersectRaySphere(r, s)
	if len(xs) != 0 {
		t.Errorf("Sphere Transform t, no intersection, got %d", len(xs))
	}
}

func TestSphereNormal(t *testing.T) {
	s := NewSphere(0, 0, 0, 1)
	n := s.NormalAt(PointV(0, 0, 1))
	if !mat.EqualApprox(n, VectorV(0, 0, 1), 0.001) {
		t.Errorf("Sphere Normal, should be (0,0,1), got %v", n)
	}
}

func TestSphereNormalTransform(t *testing.T) {
	s := NewSphere(0, 0, 0, 1)
	s.SetTransform(TranslateM(0, 1, 0))
	n := s.NormalAt(PointV(0, 1.70711, -0.70711))
	if !mat.EqualApprox(n, VectorV(0, 0.70711, -0.70711), 0.001) {
		t.Errorf("Sphere Normal, should be (0,0.70711,-0.70711), got %v", n)
	}

}
