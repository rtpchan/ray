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

func TestWorldPrepareComputation(t *testing.T) {
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

func TestWorldPrepareComputation1(t *testing.T) {
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

func TestWorldColourAt(t *testing.T) {
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

	outer := w.Object[0]
	outer.GetMaterial().Ambient = 1
	inner := w.Object[1]
	inner.GetMaterial().Ambient = 1
	r4 := NewRayCoor(0, 0, 0.75, 0, 0, -1)
	c4 := w.ColourAt(r4)
	if !ColourApprox(c4, inner.GetMaterial().Colour, 0.001) {
		t.Errorf("Get Colour, should get inner colour %v, got %v", inner.GetMaterial().Colour, c4)
	}

}

func TestWorldTransform(t *testing.T) {
	from := PointV(0, 0, 0)
	to := PointV(0, 0, -1)
	up := VectorV(0, 1, 0)
	tr := ViewTransform(from, to, up)
	if !mat.EqualApprox(tr, Identity(), 0.0001) {
		t.Errorf("World Transform, should get Identity, got %v", tr)
	}

	to = PointV(0, 0, 1)
	tr = ViewTransform(from, to, up)
	if !mat.EqualApprox(tr, ScaleM(-1, 1, -1), 0.0001) {
		t.Errorf("World Transform, should get %v, got %v", ScaleM(-1, 1, -1), tr)
	}

	from = PointV(0, 0, 8)
	to = PointV(0, 0, 0)
	tr = ViewTransform(from, to, up)
	if !mat.EqualApprox(tr, TranslateM(0, 0, -8), 0.0001) {
		t.Errorf("World Transform, should get %v, got %v", TranslateM(0, 0, -8), tr)
	}

	from = PointV(1, 3, 2)
	to = PointV(4, -2, 8)
	up = VectorV(1, 1, 0)
	tr = ViewTransform(from, to, up)
	re := mat.NewDense(4, 4, []float64{
		-0.50709, 0.50709, 0.67612, -2.36643,
		0.76772, 0.60609, 0.12122, -2.82843,
		-0.35857, 0.59761, -0.71714, 0.00000,
		0.00000, 0.0, 0.0, 1.0,
	})
	if !mat.EqualApprox(tr, re, 0.0001) {
		t.Errorf("World Transform, should get %v, got %v", re, tr)
	}
}

func TestWorldIsShadow(t *testing.T) {
	w := NewDefaultWorld()
	p := PointV(0, 10, 0)
	s := w.IsShadow(p, w.Light[0])
	if s {
		t.Errorf("Test Shadow 1, should be false, got %t", s)
	}

	p = PointV(10, -10, 10)
	s = w.IsShadow(p, w.Light[0])
	if !s {
		t.Errorf("Test Shadow 2, should be true, got %t", s)
	}

	p = PointV(-20, 20, -20)
	s = w.IsShadow(p, w.Light[0])
	if s {
		t.Errorf("Test Shadow 3, should be false, got %t", s)
	}

	p = PointV(-2, 2, -2)
	s = w.IsShadow(p, w.Light[0])
	if s {
		t.Errorf("Test Shadow 4, should be false, got %t", s)
	}
}

func TestWorldShadowHit(t *testing.T) {
	w := NewWorld()
	w.Light = append(w.Light, NewPointLight(PointV(0, 0, -10), NewColour(1, 1, 1)))
	s1 := NewSphere(0, 0, 0, 1)
	w.Object = append(w.Object, s1)
	s2 := NewSphere(0, 0, 0, 1)
	s2.Transform = TranslateM(0, 0, 10)
	w.Object = append(w.Object, s2)

	r := NewRay(PointV(0, 0, 5), VectorV(0, 0, 1))
	i := Intersection{4, s2}
	comps := PrepareComputation(i, r)
	c := w.ShadeHit(comps)
	if !ColourApprox(c, NewColour(0.1, 0.1, 0.1), 0.001) {
		t.Errorf("World Shadow Hit, colour should be (0.1,0.1,0.1), got %v", c)
	}

	r = NewRay(PointV(0, 0, -5), VectorV(0, 0, 1))
	s := NewSphere(0, 0, 0, 1)
	s.Transform = TranslateM(0, 0, 1)
	i = Intersection{5, s}
	comps = PrepareComputation(i, r)
	if comps.OverPoint.AtVec(2) >= (-EPSILON / 2.) {
		t.Errorf("Over point should be greater than neg half EPSILON, got %v",
			comps.OverPoint.AtVec(2))
	}
	if comps.OverPoint.AtVec(2) > comps.Point.AtVec(2) {
		t.Errorf("Over point should be less than Point, got %t", false)

	}
}
