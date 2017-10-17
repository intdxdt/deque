package math

//Mid2D computes the midean coordinates
func Mid2D(a, b []float64) []float64 {
    return []float64{Mid(a[x], b[x]), Mid(a[y], b[y]), }
}

//Mid computes the mean of two values
func Mid(x, y float64) float64 {
    return (x + y) / 2.0
}
