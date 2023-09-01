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
