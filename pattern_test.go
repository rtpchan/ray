package main

import "testing"

func TestPatternTransform(t *testing.T) {
	s := NewSphere()
	s.Transform = ScaleM(2, 2, 2)
	p := NewStripePattern(White(), Black())
	p.Transform = TranslateM(0.5, 0, 0)
	c := p.PatternAtShape(s, PointV(2.5, 0, 0))

	if !ColourApprox(c, NewColour(1, 1, 1), EPSILON) {
		t.Errorf("Lighting pattern, should be (1,1,1), got %v", c)
	}
}

func TestPattern(t *testing.T) {
	p1 := NewGradientPattern(White(), Black())
	c1 := p1.PatternAt(PointV(0.25, 0, 0))
	if !ColourApprox(c1, NewColour(0.75, 0.75, 0.75), EPSILON) {
		t.Errorf("Lighting pattern, should be (0.75,0.75,0.75), got %v", c1)
	}
	p2 := NewRingPattern(White(), Black())
	c2 := p2.PatternAt(PointV(0.708, 0, 0.708))
	if !ColourApprox(c2, NewColour(0, 0, 0), EPSILON) {
		t.Errorf("Lighting pattern, should be (0,0,0), got %v", c2)
	}

	p3 := NewCheckerPattern(White(), Black())
	c3 := p3.PatternAt(PointV(0, 0, 0))
	if !ColourApprox(c3, White(), EPSILON) {
		t.Errorf("Checker pattern, should be White, got %v", c3)
	}
	c3 = p3.PatternAt(PointV(0.99, 0, 0))
	if !ColourApprox(c3, White(), EPSILON) {
		t.Errorf("Checker pattern, should be White, got %v", c3)
	}
	c3 = p3.PatternAt(PointV(1.01, 0, 0))
	if !ColourApprox(c3, Black(), EPSILON) {
		t.Errorf("Checker pattern, should be Black, got %v", c3)
	}
	c3 = p3.PatternAt(PointV(0, 0.99, 0))
	if !ColourApprox(c3, White(), EPSILON) {
		t.Errorf("Checker pattern, should be White, got %v", c3)
	}
	c3 = p3.PatternAt(PointV(0, 1.01, 0))
	if !ColourApprox(c3, Black(), EPSILON) {
		t.Errorf("Checker pattern, should be Black, got %v", c3)
	}
	c3 = p3.PatternAt(PointV(0, 0, 0.99))
	if !ColourApprox(c3, White(), EPSILON) {
		t.Errorf("Checker pattern, should be White, got %v", c3)
	}
	c3 = p3.PatternAt(PointV(0, 0, 1.01))
	if !ColourApprox(c3, Black(), EPSILON) {
		t.Errorf("Checker pattern, should be Black, got %v", c3)
	}
}
