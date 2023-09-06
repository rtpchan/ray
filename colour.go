package main

import (
	"fmt"
	"math"
)

type Colour struct {
	c [3]float64
}

func Black() Colour {
	return NewColour(0, 0, 0)
}

func White() Colour {
	return NewColour(1, 1, 1)
}

func NewColour(r, g, b float64) Colour {
	return Colour{c: [3]float64{r, g, b}}
}

func NewColourWithSlice(c [3]float64) Colour {
	return Colour{c: c}
}

func AddC(a, b Colour) Colour {
	return NewColourWithSlice(Add(a.c, b.c))
}

func SubC(a, b Colour) Colour {
	return NewColourWithSlice(Sub(a.c, b.c))
}

func ScaleC(a Colour, s float64) Colour {
	return NewColourWithSlice(Scale(a.c, s))
}

func MulC(a, b Colour) Colour {
	return NewColour(a.c[0]*b.c[0], a.c[1]*b.c[1], a.c[2]*b.c[2])
}

func (c Colour) R() float64 {
	return c.c[0]
}

func (c Colour) G() float64 {
	return c.c[1]
}

func (c Colour) B() float64 {
	return c.c[2]
}

func (c Colour) String() string {
	return fmt.Sprintf("R %0.5f G %0.5f B %0.5f", c.c[0], c.c[1], c.c[2])
}

// clamp RGB8 colour to 0 - 255
func ClampRGB8(c int) int {
	if c > 255 {
		c = 255
	}
	if c < 0 {
		c = 0
	}
	return c
}

// convert to a slice of RGB 8bit integer
func (c Colour) RGB8() [3]int {
	return [3]int{
		ClampRGB8(int(math.Round(c.R() * 255))),
		ClampRGB8(int(math.Round(c.G() * 255))),
		ClampRGB8(int(math.Round(c.B() * 255)))}
}

func (c Colour) RGB8String() string {
	return fmt.Sprintf("%d %d %d ",
		ClampRGB8(int(math.Round(c.R()*255))),
		ClampRGB8(int(math.Round(c.G()*255))),
		ClampRGB8(int(math.Round(c.B()*255))))
}

func ColourApprox(c, d Colour, tol float64) bool {
	deltaR := math.Abs(c.c[0] - d.c[0])
	deltaG := math.Abs(c.c[1] - d.c[1])
	deltaB := math.Abs(c.c[2] - d.c[2])
	return deltaR < tol && deltaG < tol && deltaB < tol
}
