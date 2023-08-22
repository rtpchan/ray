package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type PointLight struct {
	Position  *mat.VecDense
	Intensity Colour
}

func NewPointLight(p *mat.VecDense, i Colour) *PointLight {
	return &PointLight{Position: p, Intensity: i}
}

// Return colour base on Material, Light, Position, vector to Eye, Normal vector at position
// Chapter 6 page 66
func Lighting(m *Material, l *PointLight, p *mat.VecDense, e *mat.VecDense, n *mat.VecDense,
	shadow bool) Colour {
	effectiveColour := MulC(m.Colour, l.Intensity)
	lightV := NormaliseV(SubV(l.Position, p))
	ambient := ScaleC(effectiveColour, m.Ambient)
	if shadow {
		return ambient
	}

	var diffuse Colour
	var specular Colour
	// to find angle between v and n, need unit vector, equipvalent to divide by magnitude
	lightDotN := DotV(lightV, n) / MagnitudeV(lightV) / MagnitudeV(n)
	if lightDotN < 0 {
		diffuse = Black()
		specular = Black()
	} else {
		diffuse = ScaleC(effectiveColour, m.Diffuse*lightDotN)
		reflectV := ReflectV(ScaleV(-1., lightV), n)
		reflectDotEye := DotV(reflectV, e) / MagnitudeV(reflectV) / MagnitudeV(e)
		if reflectDotEye <= 0 {
			specular = Black()
		} else {
			factor := math.Pow(reflectDotEye, m.Shininess)
			// factor := 0.0
			specular = ScaleC(l.Intensity, m.Specular*factor)
		}
	}
	return AddC(ambient, AddC(diffuse, specular))
}
