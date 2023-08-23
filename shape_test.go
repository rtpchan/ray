package main

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestSphereTransform(t *testing.T) {
	r := NewRayCoor(0, 0, -5, 0, 0, 1)
	s := NewSphere()
	m := ScaleM(2, 2, 2)
	s.SetTransform(m)
	xs := s.Intersect(r)
	if xs[0].T != 3 {
		t.Errorf("Sphere Transform t should be 3, got %0.2f", xs[0].T)
	}
	if xs[1].T != 7 {
		t.Errorf("Sphere Transform t should be 7, got %0.2f", xs[1].T)
	}

	s = NewSphere()
	m = TranslateM(5, 0, 0)
	s.SetTransform(m)
	xs = s.Intersect(r)
	if len(xs) != 0 {
		t.Errorf("Sphere Transform t, no intersection, got %d", len(xs))
	}
}

func TestSphereNormal(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(PointV(0, 0, 1))
	if !mat.EqualApprox(n, VectorV(0, 0, 1), 0.001) {
		t.Errorf("Sphere Normal, should be (0,0,1), got %v", n)
	}
}

func TestSphereNormalTransform(t *testing.T) {
	s := NewSphere()
	s.SetTransform(TranslateM(0, 1, 0))
	n := s.NormalAt(PointV(0, 1.70711, -0.70711))
	if !mat.EqualApprox(n, VectorV(0, 0.70711, -0.70711), 0.001) {
		t.Errorf("Sphere Normal, should be (0,0.70711,-0.70711), got %v", n)
	}
}

func TestShapePlane(t *testing.T) {
	p := NewPlane()
	n := p.LocalNormalAt(PointV(10, 0, -10))
	if !mat.EqualApprox(n, VectorV(0, 1, 0), EPSILON) {
		t.Errorf("Plane normal should be (0,1,0), got %v", n)
	}
	n = p.LocalNormalAt(PointV(0, 0, 0))
	if !mat.EqualApprox(n, VectorV(0, 1, 0), EPSILON) {
		t.Errorf("Plane normal should be (0,1,0), got %v", n)
	}

	r := NewRayCoor(0, 10, 0, 0, 0, 1)
	xs := p.Intersect(r)
	if len(xs) != 0 {
		t.Errorf("Plane - ray has no intersect, got %d intersect", len(xs))
	}
	r = NewRayCoor(0, 0, 0, 0, 0, 1)
	xs = p.Intersect(r)
	if len(xs) != 0 {
		t.Errorf("Plane - ray has no intersect, got %d intersect", len(xs))
	}

	r = NewRayCoor(0, 1, 0, 0, -1, 0)
	xs = p.Intersect(r)
	if len(xs) != 1 {
		t.Errorf("Plane - ray has 1 intersection, got %d", len(xs))
	}
	if xs[0].T != 1 {
		t.Errorf("Plane - ray T is  1 , got %f", xs[0].T)
	}
	r = NewRayCoor(0, -1, 0, 0, 1, 0)
	xs = p.Intersect(r)
	if len(xs) != 1 {
		t.Errorf("Plane - ray has 1 intersection, got %d", len(xs))
	}
	if xs[0].T != 1 {
		t.Errorf("Plane - ray T is  1 , got %f", xs[0].T)
	}
}
