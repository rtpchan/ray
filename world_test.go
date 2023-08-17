package main

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestDefaultWorld(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRayCoor(0, 0, -5, 0, 0, 1)
	xs := w.Intersect(r)
	if xs.Count() != 4 {
		t.Errorf("Test Default World, 4 intersections, got %d", xs.Count())
	} else {
		if xs[0].T != 4 {
			t.Errorf("Test Default World, [0].T=4, got %0.2f", xs[0].T)
		}
		if xs[1].T != 4.5 {
			t.Errorf("Test Default World, [0].T=4.5, got %0.2f", xs[1].T)
		}
		if xs[2].T != 5.5 {
			t.Errorf("Test Default World, [0].T=5.5, got %0.2f", xs[2].T)
		}
		if xs[3].T != 6. {
			t.Errorf("Test Default World, [0].T=6, got %0.2f", xs[3].T)
		}
	}
}

func TestPrepareComputation(t *testing.T) {
	r := NewRayCoor(0, 0, -5, 0, 0, 1)
	s := NewSphere(0, 0, 0, 1)
	i := Intersection{T: 4, Object: s}
	comps := PrepareComputation(i, r)
	if comps.T != i.T {
		t.Errorf("Prepare Computation, expect %0.2f, got %0.2f", i.T, comps.T)
	}
	if comps.Object != i.Object {
		t.Errorf("Prepare Computation, expect sphere, got %v", comps.Object)
	}
	if !mat.EqualApprox(comps.Point, PointV(0, 0, -1), 0.0001) {
		t.Errorf("Prepare Computation, expect (0,0,-1), got %v", comps.Point)
	}
	if !mat.EqualApprox(comps.EyeV, VectorV(0, 0, -1), 0.0001) {
		t.Errorf("Prepare Computation, expect (0,0,-1), got %v", comps.EyeV)
	}
	if !mat.EqualApprox(comps.NormalV, VectorV(0, 0, -1), 0.0001) {
		t.Errorf("Prepare Computation, expect (0,0,-1), got %v", comps.NormalV)
	}
}

func TestPrepareComputation1(t *testing.T) {
	r := NewRayCoor(0, 0, 0, 0, 0, 1)
	s := NewSphere(0, 0, 0, 1)
	i := Intersection{T: 1, Object: s}
	comps := PrepareComputation(i, r)
	if !mat.EqualApprox(comps.Point, PointV(0, 0, 1), 0.0001) {
		t.Errorf("Prepare Computation, expect 0,0,1, got %v", comps.Point)
	}
	if comps.Inside != true {
		t.Errorf("Prepare Computation, expect inside, got %t", comps.Inside)
	}
	if !mat.EqualApprox(comps.EyeV, VectorV(0, 0, -1), 0.0001) {
		t.Errorf("Prepare Computation, expect (0,0,-1), got %v", comps.EyeV)
	}
	if !mat.EqualApprox(comps.NormalV, VectorV(0, 0, -1), 0.0001) {
		t.Errorf("Prepare Computation, expect (0,0,-1), got %v", comps.NormalV)
	}

}

func TestColourAt(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRayCoor(0, 0, -5, 0, 1, 0)
	c := w.ColourAt(r)
	if !ColourApprox(c, NewColour(0, 0, 0), 0.0001) {
		t.Errorf("Get Colour in World, should be Black, got %v", c)
	}
	r2 := NewRayCoor(0, 0, -5, 0, 0, 1)
	c2 := w.ColourAt(r2)
	c3 := NewColour(0.38066, 0.47583, 0.2855)
	if !ColourApprox(c2, c3, 0.0001) {
		t.Errorf("Get Colour in World, should be %v, got %v", c3, c2)
	}
}
