package main

import (
	"fmt"
	"os"
)

type Canvas struct {
	c    []Colour
	w, h int
}

func NewCanvas(w, h int) *Canvas {
	return &Canvas{c: make([]Colour, w*h), w: w, h: h}
}

// covert coordinate i, j into internal slice index
func ToIndex(cv *Canvas, i, j int) int {
	if i >= cv.w {
		panic(fmt.Sprintf("Index %d exceeded canvas width %d", i, cv.w))
	}
	if j >= cv.h {
		panic(fmt.Sprintf("Index %d exceeded canvas width %d", j, cv.h))
	}
	return i + j*cv.w
}

// covert  internal slice index into coordinate i, j
func ToIJ(cv *Canvas, i int) (int, int) {
	if i >= cv.w*cv.h {
		panic(fmt.Sprintf("Index %d exceeded pixel count %d", i, cv.w*cv.h))
	}
	y := i / cv.w
	b := y * cv.w
	rem := i - b
	return rem, y
}

// Colour at pixel i, j
func (cv *Canvas) At(i, j int) Colour {
	index := ToIndex(cv, i, j)
	return cv.c[index]
}

// Write colour to pixel
func (cv *Canvas) Write(c Colour, i, j int) {
	index := ToIndex(cv, i, j)
	cv.c[index] = c
}

// Write canvas to PPM string
func (cv *Canvas) WritePPMString() string {
	s := fmt.Sprintf("P3\n%d %d\n255\n", cv.w, cv.h)
	line := ""
	for _, c := range cv.c {
		cstr := c.RGB8String()
		if len(line)+len(cstr) > 70 {
			s = s + line + "\n"
			line = cstr
		} else {
			line = line + cstr
		}
		// fmt.Println(i, line)
	}
	s = s + line + "\n"
	return s
}

func (cv *Canvas) WritePPM(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot create file: %s", err))
	}
	defer f.Close()

	s := cv.WritePPMString()
	_, err = f.WriteString(s)
	if err != nil {
		panic(fmt.Sprintf("Cannot write to file: %s", err))
	}
	err = f.Sync()
	if err != nil {
		panic(fmt.Sprintf("Error writing to disk: %s", err))
	}

}
