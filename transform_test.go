package main

import (
	"math"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestChain(t *testing.T) {
	p := PointM(1, 0, 1)
	a := RotateXM(math.Pi / 2.0)
	b := ScaleM(5, 5, 5)
	c := TranslateM(10, 5, 7)
	ia := ZeroMatrix()
	ia.Mul(b, a)
	ib := ZeroMatrix()
	ib.Mul(c, ia)
	tt := ZeroVector()
	tt.MulVec(ib, p)
	r := PointM(15, 0, 7)
	if !mat.EqualApprox(tt, r, 0.00000001) {
		t.Errorf("Transformation Chaining, got %v", r)
	}

	q := AT(c, AT(b, AT(a, p)))
	if !mat.EqualApprox(q, r, 0.000001) {
		t.Errorf("Transformation Chaining, got %v", r)
	}
}
