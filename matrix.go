package main

import "gonum.org/v1/gonum/mat"

func Identity() *mat.Dense {
	a := ZeroMatrix()
	a.Set(0, 0, 1)
	a.Set(1, 1, 1)
	a.Set(2, 2, 1)
	a.Set(3, 3, 1)
	return a
}

func InverseM(m *mat.Dense) *mat.Dense {
	invM := ZeroMatrix()
	err := invM.Inverse(m)
	if err != nil {
		return Identity()
	}
	return invM
}

// 4x4 matrix filled with zero
func ZeroMatrix() *mat.Dense {
	return mat.NewDense(4, 4, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

// Column Vector filled with zero
func ZeroVector() *mat.VecDense {
	return mat.NewVecDense(4, []float64{0, 0, 0, 0})
}

func MulM(m, n *mat.Dense) *mat.Dense {
	mulM := ZeroMatrix()
	mulM.Mul(m, n)
	return mulM
}
