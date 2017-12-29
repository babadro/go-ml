package main

import (
	"fmt"
	"log"
	"math"

	"github.com/gonum/matrix/mat64"
)

func main() {
	// Create a new matrix a.
	a := mat64.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})
	b := mat64.NewDense(3, 3, []float64{8, 9, 10, 1, 4, 2, 9, 0, 2})

	// Create a third matrix of a different size.
	c := mat64.NewDense(3, 2, []float64{3, 2, 1, 4, 0, 8})

	// Add a and b.
	d := mat64.NewDense(0, 0, nil)
	d.Add(a, b)
	fd := mat64.Formatted(d, mat64.Prefix("      "))
	fmt.Printf("d = a + b = %0.4v\n\n", fd)

	// Multiply a and c.
	f := mat64.NewDense(0, 0, nil)
	f.Mul(a, c)
	ff := mat64.Formatted(f, mat64.Prefix("      "))
	fmt.Printf("f = a c = %0.4v\n\n", ff)

	// Raising a matrix to a power.
	g := mat64.NewDense(0, 0, nil)
	g.Pow(a, 5)
	fg := mat64.Formatted(g, mat64.Prefix("      "))
	fmt.Printf("g = a^5 = %0.4v\n\n", fg)

	// Apply a function to each of the elements of a.
	h := mat64.NewDense(0, 0, nil)
	sqrt := func(_, _ int, v float64) float64 { return math.Sqrt(v) }
	h.Apply(sqrt, a)
	fh := mat64.Formatted(h, mat64.Prefix("      "))
	fmt.Printf("h = sqrt(a) = %0.4v\n\n", fh)

	// Compute and output the transpose of the matrix.
	ft := mat64.Formatted(a.T(), mat64.Prefix("      "))
	fmt.Printf("a^T = %v\n\n", ft)

	// Compute and output the determinant of a.
	deta := mat64.Det(a)
	fmt.Printf("det(a) = %.2f\n\n", deta)

	// Compute and ouput the inverse of a.
	aInverse := mat64.NewDense(0, 0, nil)
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fi := mat64.Formatted(aInverse, mat64.Prefix("      "))
	fmt.Printf("a^-1 = %v\n\n", fi)
}
