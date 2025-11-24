package main

import (
	"fmt"
	"math"
)

// ====================================== Helper functions

// allows for interface types to be converted into numbers to perform operations on
func ToNumber(num PSConstant) (float64, error) {
	switch val := num.(type) {
	case int:
		return float64(val), nil
	case float64:
		return val, nil
	default:
		return math.NaN(), fmt.Errorf("incorrect input")
	}
}
