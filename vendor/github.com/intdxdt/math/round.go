package math

import "math"

//RoundFloor rounds a float to the nearest whole number float
func RoundFloor(f float64) float64 {
    return math.Trunc(f + math.Copysign(0.5, f))
}

//Round rounds a number to the nearest decimal place
func Round(x float64, digits... int) float64 {
    d := 0
    if len(digits) > d {
        d = digits[d]
    }
    m := math.Pow(10.0, float64(d))
    return RoundFloor(x *m) / m
}
