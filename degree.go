package degrees

import "math"

func d2r(dg float64) float64 {
	return dg / 180 * math.Pi
}

func Cos(dg float64) float64 {
	return math.Cos(d2r(dg))
}

func Sin(dg float64) float64 {
	return math.Sin(d2r(dg))
}

func Tan(dg float64) float64 {
	return math.Tan(d2r(dg))
}

