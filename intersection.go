package main

import "math"

// Ray report Intersection hit
type Intersection struct {
	T      float64
	Object Shape
}

type Intersections []Intersection

func (is Intersections) Len() int {
	return len(is)
}

func (is Intersections) Count() int {
	return len(is)
}

// to satisfy sort.Sort() interface
func (is Intersections) Less(i, j int) bool {
	return is[i].T < is[j].T
}

// to satisfy sort.Sort() interface
func (is Intersections) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is Intersections) Hit() Intersections {
	var hit Intersection
	t := math.MaxFloat64
	for _, i := range is {
		if i.T > 0 && i.T < t {
			t = i.T
			hit = i
		}
	}
	if t < math.MaxFloat64 {
		return []Intersection{hit}
	}
	return []Intersection{}
}
