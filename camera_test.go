package main

import (
	"math"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestCameraPixelSize(t *testing.T) {
	c := NewCamera(200, 125, math.Pi/2.)
	if c.PixelSize != 0.01 {
		t.Errorf("Camera Pixel size, should be 0.01, got %f", c.PixelSize)
	}
	c = NewCamera(125, 200, math.Pi/2.)
	if c.PixelSize != 0.01 {
		t.Errorf("Camera Pixel size, should be 0.01, got %f", c.PixelSize)
	}
}

func TestCameraRay(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2.0)
	r := c.Ray(100, 50)
	if !mat.EqualApprox(r.Origin, PointV(0, 0, 0), 0.001) {
		t.Errorf("Camera Ray to Pixel, origin shoulbe be (0,0,0), got %v", r.Origin)
	}
	if !mat.EqualApprox(r.Dir, VectorV(0, 0, -1), 0.001) {
		t.Errorf("Camera Ray to Pixel, direction shoulbe be (0, 0, -1), got %v", r.Dir)
	}

	c.Transform = CT(RotateYM(math.Pi/4.), TranslateM(0, -2., 5.))
	r = c.Ray(100, 50)
	if !mat.EqualApprox(r.Origin, PointV(0, 2, -5), 0.001) {
		t.Errorf("Camera Ray to Pixel, origin shoulbe be (0,2,-5), got %v", r.Origin)
	}
	if !mat.EqualApprox(r.Dir, VectorV(math.Sqrt2/2., 0, -math.Sqrt2/2.), 0.001) {
		t.Errorf("Camera Ray to Pixel, direction shoulbe be (root2/2, 0, -root2/2), got %v", r.Dir)
	}
}

func TestCameraRender(t *testing.T) {
	w := NewDefaultWorld()
	c := NewCamera(11, 11, math.Pi/2.0)
	from := PointV(0, 0, -5)
	to := PointV(0, 0, 0)
	up := VectorV(0, 1, 0)
	c.Transform = ViewTransform(from, to, up)
	canvas := c.Render(w)
	if !ColourApprox(canvas.At(5, 5), NewColour(0.38066, 0.47583, 0.2855), 0.001) {
		t.Errorf("Camera Render Colour, got %v", canvas.At(5, 5))
	}
}
