package darts

import "math"

func Score(x, y float64) int {
	distance := math.Sqrt(x*x + y*y)

	switch {
	case distance > 10:
		return 0
	case distance > 5:
		return 1
	case distance > 1:
		return 5
	default:
		return 10
	}
}
