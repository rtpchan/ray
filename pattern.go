package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Pattern interface {
	GetTransform() *mat.Dense
	PatternAt(*mat.VecDense) Colour
	PatternAtShape(Shaper, *mat.VecDense) Colour
}

type PatternTransform struct {
	Transform *mat.Dense
}

func NewPatternTransform() PatternTransform {
	return PatternTransform{Transform: Identity()}
}

func (pt PatternTransform) GetTransform() *mat.Dense {
	return pt.Transform
}

type StripePattern struct {
	A Colour
	B Colour
	PatternTransform
}

func NewStripePattern(a Colour, b Colour) *StripePattern {
	return &StripePattern{A: a, B: b, PatternTransform: NewPatternTransform()}
}

func (p *StripePattern) PatternAt(v *mat.VecDense) Colour {
	if math.Mod(math.Floor(v.AtVec(0)), 2) == 0 {
		return p.A
	}
	return p.B
}

func (p *StripePattern) PatternAtShape(s Shaper, v *mat.VecDense) Colour {
	objectPoint := MulV(InverseM(s.GetTransform()), v)
	patternPoint := MulV(InverseM(p.Transform), objectPoint)
	return p.PatternAt(patternPoint)
}
