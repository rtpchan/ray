package main

import (
	"math"
	"testing"
)

func TestLight(t *testing.T) {
	m := NewMaterial()
	position := PointV(0, 0, 0)
	eyeV := VectorV(0, 0, -1)
	normalV := VectorV(0, 0, -1)
	light := NewPointLight(PointV(0, 0, -10), NewColour(1, 1, 1))
	result := Lighting(m, light, position, eyeV, normalV, false)
	if !ColourApprox(result, NewColour(1.9, 1.9, 1.9), 0.0001) {
		t.Errorf("Lighting, colour should be (1.9,1.9,1.9), got %v", result)
	}

	eyeV = VectorV(0, math.Sqrt2/2, -math.Sqrt2/2)
	result = Lighting(m, light, position, eyeV, normalV, false)
	if !ColourApprox(result, NewColour(1., 1., 1.), 0.0001) {
		t.Errorf("Lighting, colour should be (1.,1.,1.), got %v", result)
	}

	eyeV = VectorV(0, 0, -1)
	light = NewPointLight(PointV(0, 10, -10), NewColour(1, 1, 1))
	result = Lighting(m, light, position, eyeV, normalV, false)
	if !ColourApprox(result, NewColour(0.7364, 0.7364, 0.7364), 0.0001) {
		t.Errorf("Lighting, colour should be (0.7364, 0.7364, 0.7364), got %v", result)
	}

	eyeV = VectorV(0, -math.Sqrt2/2, -math.Sqrt2/2)
	result = Lighting(m, light, position, eyeV, normalV, false)
	if !ColourApprox(result, NewColour(1.6364, 1.6364, 1.6364), 0.0001) {
		t.Errorf("Lighting, colour should be (1.6364,1.6364,1.6364), got %v", result)
	}

	eyeV = VectorV(0, 0, -1)
	light = NewPointLight(PointV(0, 0, 10), NewColour(1, 1, 1))
	result = Lighting(m, light, position, eyeV, normalV, false)
	if !ColourApprox(result, NewColour(0.1, 0.1, 0.1), 0.0001) {
		t.Errorf("Lighting, colour should be (0.1,0.1,0.1), got %v", result)
	}
}

func TestLightShadow(t *testing.T) {
	m := NewMaterial()
	position := PointV(0, 0, 0)
	eyeV := VectorV(0, 0, -1)
	normalV := VectorV(0, 0, -1)
	light := NewPointLight(PointV(0, 0, -10), NewColour(1, 1, 1))
	inShadow := true
	result := Lighting(m, light, position, eyeV, normalV, inShadow)

	if !ColourApprox(result, NewColour(0.1, 0.1, 0.1), 0.0001) {
		t.Errorf("Lighting in shadow, should be (0.1,0.1,0.1), got %v", result)
	}
}
