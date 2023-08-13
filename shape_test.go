package main

import "testing"

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
