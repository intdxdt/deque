package math

import "github.com/intdxdt/robust"

//Det2 computes the determinant of 2x2 matrix
func Det2(mat2x2 [][]float64) float64 {
	res := robust.Det2(mat2x2)
	return res[len(res)-1]
}

//Det3 computes the determinant of a 3x3 matrix
func Det3(mat3x3 [][]float64) float64 {
	res := robust.Det3(mat3x3)
	return res[len(res)-1]
}
