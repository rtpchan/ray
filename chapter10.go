package main

import "math"

func Chapter10() {

	exportFilename := "chapter10.ppm"

	plane := NewPlane()
	plane.Material.Colour = NewColour(1, 0.9, 0.9) // Solid Colour ignored
	plane.Material.Pattern = NewCheckerPattern(NewColour(1, 0, 0), NewColour(0, 1, 0))
	plane.Material.Specular = 0

	// floor := NewSphere()
	// floor.Transform = ScaleM(10, 0.01, 10)
	// floor.Material = NewMaterial()
	// floor.Material.Colour = NewColour(1, 0.9, 0.9)
	// floor.Material.Specular = 0

	// leftWall := NewSphere()
	// leftWall.Transform = CT(TranslateM(0, 0, 5), CT(RotateYM(-math.Pi/4.),
	// 	CT(RotateXM(math.Pi/2.), ScaleM(10, 0.01, 10))))
	// leftWall.Material = floor.Material

	// rightWall := NewSphere()
	// rightWall.Transform = CT(TranslateM(0, 0, 5), CT(RotateYM(math.Pi/4.),
	// 	CT(RotateXM(math.Pi/2.), ScaleM(10, 0.01, 10))))
	// rightWall.Material = floor.Material

	middle := NewSphere()
	middle.Transform = TranslateM(-0.5, 1, 3.5)
	middle.Material = NewMaterial()
	middle.Material.Colour = NewColour(0.1, 1, 0.5)
	mp := NewRingPattern(NewColour(0, 1, 1), NewColour(1, 1, 0))
	mp.Transform = ScaleM(0.3, 0.3, 0.3)
	middle.Material.Pattern = mp
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3

	right := NewSphere()
	right.Transform = CT(TranslateM(1.5, 0.5, -0.5), ScaleM(0.5, 0.5, 0.5))
	right.Material = NewMaterial()
	right.Material.Colour = NewColour(0.5, 1, 0.1)
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.3

	left := NewSphere()
	left.Transform = CT(TranslateM(-1.5, 0.33, -0.75), ScaleM(0.33, 0.33, 0.33))
	left.Material = NewMaterial()
	left.Material.Colour = NewColour(1, 0.8, 0.1)
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3

	w := NewWorld()
	w.Light = append(w.Light, NewPointLight(PointV(-10, 10, -10), NewColour(1, 1, 1)))
	w.Object = []Shaper{plane, middle, left, right}

	c := NewCamera(300, 150, math.Pi/3.)
	// c.Transform = ViewTransform(PointV(0, 1.5, -5), PointV(0, 1, 0), VectorV(0, 1, 0))
	c.Transform = ViewTransform(PointV(0, 1.5, -5), PointV(0, 1, 0), VectorV(0, 1, 0))
	canvas := c.Render(w)
	canvas.WritePPM(exportFilename)
}
