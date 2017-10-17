package math

import "math"

// Min returns the smaller of x or y.
//
// Special cases are:
//	Min(x, -Inf) = Min(-Inf, x) = -Inf
//	Min(x, NaN) = Min(NaN, x) = NaN
//	Min(-0, ±0) = Min(±0, -0) = -0
func MinF64(x, y float64) float64{
	return math.Min(x, y)
}

// Max returns the larger of x or y.
//
// Special cases are:
//	Max(x, +Inf) = Max(+Inf, x) = +Inf
//	Max(x, NaN) = Max(NaN, x) = NaN
//	Max(+0, ±0) = Max(±0, +0) = +0
//	Max(-0, -0) = -0
func MaxF64(x, y float64) float64{
	return math.Max(x, y)
}

func MinI64(x, y int64) int64 {
    if y < x {
        return y
    }
    return x
}

func MaxI64(x, y int64) int64 {
    if y > x {
        return y
    }
    return x
}

func MinInt(x, y int) int {
    if y < x {
        return y
    }
    return x
}

func MaxInt(x, y int) int {
    if y > x {
        return y
    }
    return x
}
