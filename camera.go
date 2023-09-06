package main

import (
	"math"
	"sync"

	"gonum.org/v1/gonum/mat"
)

type Camera struct {
	Hsize     int
	Vsize     int
	FOV       float64
	Transform *mat.Dense

	PixelSize  float64
	HalfWidth  float64
	HalfHeight float64
}

func NewCamera(h, v int, fov float64) *Camera {
	c := &Camera{Hsize: h, Vsize: v, FOV: fov, Transform: Identity()}
	c.pixelSize()
	return c
}

func (c *Camera) pixelSize() {
	halfView := math.Tan(c.FOV / 2.0)
	aspect := float64(c.Hsize) / float64(c.Vsize)
	if aspect >= 1.0 {
		c.HalfWidth = halfView
		c.HalfHeight = halfView / aspect
	} else {
		c.HalfWidth = halfView * aspect
		c.HalfHeight = halfView
	}
	c.PixelSize = (c.HalfWidth * 2.) / float64(c.Hsize)
}

func (c *Camera) Ray(px, py int) *Ray {
	xOffset := (float64(px) + 0.5) * c.PixelSize
	yOffset := (float64(py) + 0.5) * c.PixelSize
	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset
	pixel := MulV(InverseM(c.Transform), PointV(worldX, worldY, -1))
	// origin := MulV(InverseM(c.Transform), PointV(0, 0, 0))
	origin := AT(InverseM(c.Transform), PointV(0, 0, 0))
	direction := NormaliseV(SubV(pixel, origin))
	return NewRay(origin, direction)

}

func (c *Camera) Render(w *World) *Canvas {
	cv := NewCanvas(c.Hsize, c.Vsize)
	var wg sync.WaitGroup
	for j := 0; j < c.Vsize; j++ {
		for i := 0; i < c.Hsize; i++ {
			wg.Add(1)
			i := i
			j := j
			go func() {
				defer wg.Done()
				c.onePixel(w, cv, i, j)
			}()
		}
	}
	wg.Wait()
	return cv
}

func (c *Camera) onePixel(w *World, cv *Canvas, i, j int) {
	ray := c.Ray(i, j)
	colour := w.ColourAt(ray, 5) // TODO hardcode maximum 5 reflection
	cv.Write(colour, i, j)
}
