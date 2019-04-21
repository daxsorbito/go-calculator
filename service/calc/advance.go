package calc

import (
	"math"
)

// Sqrt method for Sqrt of an int64 argument
func Sqrt(arg int64) float32 {
	return float32(math.Sqrt(float64(arg)))
}

// Cbrt method for Cbrt of an int64 argument
func Cbrt(arg int64) float32 {
	return float32(math.Cbrt(float64(arg)))
}

// Fac method for Factorial of an int64 argument
func Fac(arg int64) float32 {
	result := float32(1)
	for x := arg; x > 0; x-- {
		result *= float32(x)
	}
	return float32(result)
}
