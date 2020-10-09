package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(3 + 2i)
)

const pi = 3.141

func main() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var x, y int = 3, 4
	var f64 float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f64)
	fmt.Println(x, y, z)
}
