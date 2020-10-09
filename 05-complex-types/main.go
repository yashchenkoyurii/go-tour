package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X int
	Y int
}

var v = &Vertex{1, 2}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

var a [2]string // Array

// The length of a slice is the number of elements it contains.
// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	fmt.Println(v.X)
	v.X = 123
	fmt.Println(v.X)

	fmt.Println(v1, p, v2, v3)

	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:3] // Slice
	fmt.Println(s)

	names := [4]string{
		"First",
		"Second",
		"Third",
		"Fourth",
	}

	s1 := names[1:3]
	s2 := names[3:4]
	s1[0] = "XXX"

	fmt.Println(s1, s2)
	fmt.Println(names)

	s3 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s3)

	s3 = s3[:0]
	printSlice(s3)

	s3 = s3[:4]
	printSlice(s3)

	s3 = s3[2:]
	printSlice(s3)

	var s4 []int
	printSlice(s4)
	if s4 == nil {
		fmt.Println("nil!")
	}

	s4 = append(s4, 0)
	printSlice(s4)
	s4 = append(s4, 1)
	printSlice(s4)
	s4 = append(s4, 2, 3, 4)
	printSlice(s4)

	var pow []int = []int{1, 2, 4, 8, 16, 32}

	for idx, value := range pow {
		fmt.Printf("2**%d=%d\n", idx, value)
	}

	//Maps
	var m map[string]Vertex

	m = make(map[string]Vertex)
	m["KEY"] = Vertex{3, 4}

	elem, ok := m["KEY"]
	fmt.Println(elem, ok)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
