package math

import "strconv"


//Deg2rad converts degrees to radians
func Deg2rad(d float64) float64 {
    return d * (Pi / 180.0)
}

//Rad2deg converts radians to degrees
func Rad2deg(r float64) float64 {
    return r * (180.0 / Pi)
}

//float number to a string
func FloatToString(v float64) string {
    return strconv.FormatFloat(v, 'f', -1, 64)
}
