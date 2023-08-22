package main

import "testing"

func TestIntersectionHit(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{T: 1, Object: s}
	i2 := Intersection{T: 2, Object: s}
	is := Intersections{i1, i2}
	if is.Count() != 2 {
		t.Errorf("Intersection, should be 2, got %d", is.Count())
	}

	ir := is.Hit()
	if len(ir) != 1 && ir[0] != i1 {
		t.Errorf("Intersection, should be i1, got %v", ir)
	}
}

func TestIntersectionHit2(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{T: -1, Object: s}
	i2 := Intersection{T: 1, Object: s}
	is := Intersections{i1, i2}

	ir := is.Hit()
	if len(ir) != 1 && ir[0] != i2 {
		t.Errorf("Intersection, should be i1, got %v", ir)
	}
}

func TestIntersectionHit3(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{T: -2, Object: s}
	i2 := Intersection{T: -1, Object: s}
	is := Intersections{i1, i2}

	ir := is.Hit()
	if len(ir) != 0 {
		t.Errorf("Intersection, should have no intersection, got %v", ir)
	}
}

func TestIntersectionHit4(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{T: 5, Object: s}
	i2 := Intersection{T: 7, Object: s}
	i3 := Intersection{T: -3, Object: s}
	i4 := Intersection{T: 2, Object: s}
	is := Intersections{i1, i2, i3, i4}

	ir := is.Hit()
	if len(ir) != 1 && ir[0] != i4 {
		t.Errorf("Intersection, should be i4, got %v", ir)
	}
}
