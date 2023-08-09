package main

import (
	"testing"
)

func TestCanvas(t *testing.T) {
	cv := NewCanvas(20, 10)

	index := ToIndex(cv, 5, 6)
	if index != 125 {
		t.Errorf("Canvas index from coor, should be 125, got %d", index)
	}

	x, y := ToIJ(cv, 65)
	if x != 5 {
		t.Errorf("Canvas coor from index, should be 5, got %d", x)
	}
	if y != 3 {
		t.Errorf("Canvas coor from index, should be 6, got %d", y)
	}

	b := cv.At(8, 9)
	if b.R() != 0.0 || b.G() != 0.0 || b.B() != 0.0 {
		t.Errorf("Pixel should be black, got %v", b.c)
	}

	cv.Write(NewColour(1.0, 0.2, 0.1), 8, 9)
	d := cv.At(8, 9)
	if d.R() != 1.0 || d.G() != 0.2 || d.B() != 0.1 {
		t.Errorf("Pixel should be R 1.0 G 0.2 B 0.1, got %v", d.c)
	}
}

func TestCanvasPPM1(t *testing.T) {
	cv := NewCanvas(5, 3)
	c1 := NewColour(1.5, 0, 0)
	c2 := NewColour(0, 0.5, 0)
	c3 := NewColour(-0.5, 0, 1)
	cv.Write(c1, 0, 0)
	cv.Write(c2, 2, 1)
	cv.Write(c3, 4, 2)
	s := cv.WritePPMString()
	output := `P3
5 3
255
255 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 128 0 0 0 0 0 0 0 0 0 0 
0 0 0 0 0 0 0 0 0 0 0 255 
`
	if s != output {
		t.Errorf("Write PPM String, should get %s, got %s", output, s)
	}
}

func TestCanvasPPM2(t *testing.T) {
	cv := NewCanvas(10, 2)
	c1 := NewColour(1, 0.8, 0.6)
	for i := 0; i < 10; i++ {
		for j := 0; j < 2; j++ {
			cv.Write(c1, i, j)
		}
	}
	s := cv.WritePPMString()
	output := `P3
10 2
255
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 
`
	if s != output {
		t.Errorf("Write PPM String, should get %s, got %s", output, s)
	}
	// cv.WritePPM("test.ppm")
}
