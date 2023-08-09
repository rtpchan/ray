package main

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestMatrixMul(t *testing.T) {
	a := mat.NewDense(4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2})
	b := mat.NewDense(4, 4, []float64{-2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 1, 2, 7, 8})
	// c := mat.NewDense(4, 4, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	c := ZeroMatrix()
	// var c *mat.Dense
	c.Mul(a, b)
	d := mat.NewDense(4, 4, []float64{20, 22, 50, 48, 44, 54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42})
	if !mat.Equal(c, d) {
		t.Errorf("Matrix multiplcation, got %v", c)
	}

	e := mat.NewDense(4, 4, []float64{1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1})
	f := mat.NewVecDense(4, []float64{1, 2, 3, 1})
	g := ZeroVector()
	g.MulVec(e, f)
	k := mat.NewVecDense(4, []float64{18, 24, 33, 1})
	if !mat.Equal(g, k) {
		t.Errorf("Matrix Vector multiplication, got %v", g)
	}

}

func TestMatrixTranspose(t *testing.T) {
	a := mat.NewDense(4, 4, []float64{0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8})
	b := mat.NewDense(4, 4, []float64{0, 9, 1, 0, 9, 8, 8, 0, 3, 0, 5, 5, 0, 8, 3, 8})
	c := a.T()
	if !mat.Equal(b, c) {
		t.Errorf("Matrix transpose, got %v", c)
	}
}

func TestMatrixDeterminant(t *testing.T) {
	a := mat.NewDense(2, 2, []float64{1, 5, -3, 2})
	d := mat.Det(a)
	if d != 17 {
		t.Errorf("Matrix determinant, should be 17, got %f", d)
	}
}

func TestMatrixInverse(t *testing.T) {
	a := mat.NewDense(4, 4, []float64{-5, 2, 6, -8, 1, -5, 1, 8, 7, 7, -6, -7, 1, -3, 7, 4})
	b := ZeroMatrix()
	err := b.Inverse(a)
	if err != nil {
		t.Errorf("Matrix inverse failed")
	}

}
