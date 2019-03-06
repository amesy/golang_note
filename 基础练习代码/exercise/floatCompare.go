package exercise

import (
	"math"
)

const MIN = 0.000001

func Compare(f1, f2 float64) bool {
	return math.Dim(f1, f2) < MIN
}


