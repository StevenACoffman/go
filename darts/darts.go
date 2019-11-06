package darts

import (
	"math"
)

func inCircle(radius int, x, y float64) bool {
	distance := math.Sqrt((x * x) + (y * y))
	return distance <= float64(radius)
}

func Score(x, y float64) int {
	if inCircle(1, x, y) {
		return 10
	} else if inCircle(5, x, y) {
		return 5
	} else if inCircle(10, x, y) {
		return 1
	}
	return 0
}
