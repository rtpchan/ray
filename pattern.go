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

type GradientPattern struct {
	A Colour
	B Colour
	PatternTransform
}

func NewGradientPattern(a Colour, b Colour) *GradientPattern {
	return &GradientPattern{A: a, B: b, PatternTransform: NewPatternTransform()}
}

func (p *GradientPattern) PatternAt(v *mat.VecDense) Colour {
	distance := SubC(p.B, p.A)
	x := v.AtVec(0)
	fraction := x - math.Floor(x)
	return AddC(p.A, ScaleC(distance, fraction))
}

func (p *GradientPattern) PatternAtShape(s Shaper, v *mat.VecDense) Colour {
	objectPoint := MulV(InverseM(s.GetTransform()), v)
	patternPoint := MulV(InverseM(p.Transform), objectPoint)
	return p.PatternAt(patternPoint)
}

type RingPattern struct {
	A Colour
	B Colour
	PatternTransform
}

func NewRingPattern(a Colour, b Colour) *RingPattern {
	return &RingPattern{A: a, B: b, PatternTransform: NewPatternTransform()}
}

func (p *RingPattern) PatternAt(v *mat.VecDense) Colour {
	x := v.AtVec(0)
	z := v.AtVec(2)
	dist := math.Sqrt(x*x + z*z)
	if math.Mod(math.Floor(dist), 2.) == 0 {
		return p.A
	}
	return p.B
}

type CheckerPattern struct {
	A Colour
	B Colour
	PatternTransform
}

func (p *RingPattern) PatternAtShape(s Shaper, v *mat.VecDense) Colour {
	objectPoint := MulV(InverseM(s.GetTransform()), v)
	patternPoint := MulV(InverseM(p.Transform), objectPoint)
	return p.PatternAt(patternPoint)
}

func NewCheckerPattern(a Colour, b Colour) *CheckerPattern {
	return &CheckerPattern{A: a, B: b, PatternTransform: NewPatternTransform()}
}

func (p *CheckerPattern) PatternAt(v *mat.VecDense) Colour {
	x := math.Floor(v.AtVec(0))
	y := math.Floor(v.AtVec(1))
	z := math.Floor(v.AtVec(2))

	if math.Mod((x+y+z), 2) == 0 {
		return p.A
	}
	return p.B
}

func (p *CheckerPattern) PatternAtShape(s Shaper, v *mat.VecDense) Colour {
	objectPoint := MulV(InverseM(s.GetTransform()), v)
	patternPoint := MulV(InverseM(p.Transform), objectPoint)
	return p.PatternAt(patternPoint)
}
