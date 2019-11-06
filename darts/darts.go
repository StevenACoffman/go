package darts

import "math"

// Score returns the score
func Score(x, y float64) int {
	switch {
	case inCircle(x, y, 1):
		return 10
	case inCircle(x, y, 5):
		return 5
	case inCircle(x, y, 10):
		return 1
	default:
		return 0
	}
}

func inCircle(x, y, radius float64) bool {
	dist := math.Sqrt(x*x + y*y)
	return dist <= radius
}
