// Package polygo is a collection of tools that make working with polynomials easier in Go.
package polygo

/*
This file contains general utility functions used throughout the library.
*/

// Strip all leading zeroes in the slice s. If the entire slice is filled with 0, the first element will be kept.
func stripTailingZeroes(s []float64) []float64 {
	for s[len(s)-1] == 0.0 && len(s) > 1 {
		s = s[:len(s)-1]
	}
	return s
}

// Cast []float64 to []complex128
func complex128Slice(s []float64) []complex128 {
	ret := make([]complex128, len(s))
	for i, e := range s {
		ret[i] = complex(e, 0)
	}
	return ret
}

// Cast []complex128 to []float64
func float64Slice(s []complex128) []float64 {
	ret := make([]float64, len(s))
	for i, e := range s {
		ret[i] = float64(real(e))
	}
	return ret
}

// Finds the smallest power of 2 greater than n.
func nextClosestPowerOfTwo(n int) int {
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n++
	return n
}
