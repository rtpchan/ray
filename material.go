package main

type Material struct {
	Colour    Colour
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64 // 10 very shinny, 200 not so shinny
}

// func NewMaterial(c Colour, a, d, s, sh float64) *Material {
// 	return &Material{Colour: c, Ambient: a, Diffuse: d, Specular: s, Shininess: sh}
// }

func NewMaterial() *Material {

	return &Material{Colour: NewColour(1, 1, 1),
		Ambient: 0.1, Diffuse: 0.9, Specular: 0.9, Shininess: 200}
}
