package darts

import (
	"math"
)

/*
The outer circle has a radius of 10 units
the middle circle a radius of 5 units,
and the inner circle a radius of 1
*/
func Score(x, y float64) (score int) {
	switch distance := math.Sqrt((x * x) + (y * y)); {
	case distance <= 1.0:
		score = 10
	case distance <= 5.0:
		score = 5
	case distance <= 10.0:
		score = 1
	default:
		score = 0
	}

	return score
}
