package polynomial_test

import (
	"fmt"
	"main/mathlib/polynomial"
	"testing"
)

var Cons1, Cons2, Line1, Line2, Quad1, Quad2, Cube1 *polynomial.RealPolynomial

func init() {
	// Testing polynomials
	var err error
	Cons1, err = polynomial.NewRealPolynomial([]float64{0})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	Cons2, err = polynomial.NewRealPolynomial([]float64{18})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	Line1, err = polynomial.NewRealPolynomial([]float64{6, 5})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	Line2, err = polynomial.NewRealPolynomial([]float64{4, 3})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	Quad1, err = polynomial.NewRealPolynomial([]float64{-1, 1, 1, 0, 0, 0, 0, 0})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	Quad2, err = polynomial.NewRealPolynomial([]float64{12, 42, -56})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	Cube1, err = polynomial.NewRealPolynomial([]float64{11, 7, 3, 5})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
}

func TestFindRootWithin(t *testing.T) {
	root, err := Cons1.FindRootWithin(-5, 5)
	Cons1.PrintExpr()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("Root: %f\n", root)
}

func TestCountRoots(t *testing.T) {
	nRoots, err := Cons1.CountRootsWithin(-5, 5)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	Cons1.PrintExpr()
	fmt.Printf("Number of roots: %d\n", nRoots)
}
