package main

import (
	"sort"

	"gonum.org/v1/gonum/mat"
)

type World struct {
	Object []Shape
	Light  []*PointLight
}

func NewWorld() *World {
	return &World{}
}

func NewDefaultWorld() *World {
	w := &World{}
	w.Light = []*PointLight{NewPointLight(PointV(-10, 10, -10), NewColour(1, 1, 1))}
	s1 := NewSphere(0, 0, 0, 1)
	m1 := NewMaterial()
	m1.Colour = NewColour(0.8, 1.0, 0.6)
	m1.Diffuse = 0.7
	m1.Specular = 0.2
	s1.Material = m1
	s2 := NewSphere(0, 0, 0, 1)
	s2.Transform = ScaleM(0.5, 0.5, 0.5)
	w.Object = []Shape{s1, s2}
	return w
}

// Find intersection by a ray
func (w *World) Intersect(r *Ray) Intersections {
	is := Intersections{}
	for _, obj := range w.Object {
		xs := IntersectRaySphere(r, obj)
		if xs.Count() > 0 {
			is = append(is, xs...)
		}
	}
	sort.Sort(is)
	return is
}

// Find colour of a hit by a ray, c is pre-computed with PrepareComputation()
func (w *World) ShadeHit(c *Comps) Colour {
	col := Black()
	for _, light := range w.Light {
		inShadow := w.IsShadow(c.OverPoint, light)
		lg := Lighting(c.Object.GetMaterial(), light,
			c.Point, c.EyeV, c.NormalV, inShadow)
		col = AddC(lg, col)
	}
	return col
}

// get resulting colour for a ray from the eye
func (w *World) ColourAt(r *Ray) Colour {
	xs := w.Intersect(r)
	for _, x := range xs {
		if x.T < 0 {
			continue
		} else {
			c := PrepareComputation(x, r)
			return w.ShadeHit(c)
		}
	}
	return Black() // return black if no hit
}

// test if Point p is in shadow from Light l
func (w *World) IsShadow(p *mat.VecDense, l *PointLight) bool {
	v := SubV(l.Position, p)
	distance := MagnitudeV(v)
	direction := NormaliseV(v)
	r := NewRay(p, direction)
	intersections := w.Intersect(r)
	h := intersections.Hit()
	if len(h) > 0 {
		return h[0].T < distance
	}
	return false
}

type Comps struct {
	T         float64
	Object    Shape
	Point     *mat.VecDense
	EyeV      *mat.VecDense
	NormalV   *mat.VecDense
	Inside    bool
	OverPoint *mat.VecDense
}

// Prepare intersection information
func PrepareComputation(i Intersection, r *Ray) *Comps {
	c := &Comps{
		T:      i.T,
		Object: i.Object,
	}
	p := r.Position(i.T)
	nv := i.Object.NormalAt(p)

	c.Point = p
	c.EyeV = ScaleV(-1, r.Dir)

	if DotV(nv, c.EyeV) < 0 { // inside object, normal point to the other side
		c.Inside = true
		c.NormalV = ScaleV(-1, nv)
	} else {
		c.Inside = false
		c.NormalV = nv
	}
	c.OverPoint = AddV(c.Point, ScaleV(EPSILON, c.NormalV))
	return c
}

// Chapter 7 page 99, transform matrix to move the world to fit the camera
func ViewTransform(from, to, up *mat.VecDense) *mat.Dense {
	forward := NormaliseV(SubV(to, from))
	left := CrossV(forward, NormaliseV(up))
	trueUp := CrossV(left, forward)
	backward := ScaleV(-1, forward)
	orientation := mat.NewDense(4, 4, []float64{
		left.AtVec(0), left.AtVec(1), left.AtVec(2), 0,
		trueUp.AtVec(0), trueUp.AtVec(1), trueUp.AtVec(2), 0,
		backward.AtVec(0), backward.AtVec(1), backward.AtVec(2), 0,
		0, 0, 0, 1,
	})
	fromNeg := TranslateM(-from.AtVec(0), -from.AtVec(1), -from.AtVec(2))
	return MulM(orientation, fromNeg)
}
