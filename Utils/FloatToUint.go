package Utils

import (
	"math"
)

func ParseFloat64ToUint(number float64) uint {
	uintValue := uint(math.Round(number))
	return uint(uintValue)
}
