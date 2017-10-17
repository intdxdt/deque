package math

import "math"

//Sign compute the 1.0 x sign of a number, return 0.0 when +/- 0.0
func Sign(n float64) float64 {
	//handles +/- 0.0
	if Abs(n) == 0 {
		return 0
	}
	return math.Copysign(1.0, n)
}

// Sign of a 2x2 determinant - robustly.
func SignOfDet2(x1, y1, x2, y2 float64) int {
	// returns -1 if the determinant is negative,
	// returns  1 if the determinant is positive,
	// retunrs  0 if the determinant is null.
	return int(Sign(Det2([][]float64{{x1, y1}, {x2, y2}})))
}
