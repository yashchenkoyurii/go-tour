package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if res := math.Pow(x, n); res < lim {
		return res
	}

	return lim
}

func main() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")

	sum := 0

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		sum += i
	}

	fmt.Println(sum)

	x := 1
	for x < 1000 {
		x += x
	}

	fmt.Println(x)
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
