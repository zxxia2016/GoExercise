package main

import (
	"errors"
	"math"
)

func main() {
	v, e := MySqrt(4)
	if e == nil {
		print(v)
	}
}

func MySqrt(a float64) (float64, error) {
	if a < 0 {
		return float64(math.NaN()), errors.New("invalid number")
	}
	return math.Sqrt(a), nil
}
