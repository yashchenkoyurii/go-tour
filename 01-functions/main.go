package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 1 / 2
	y = sum - x

	return
}

func main() {
	fmt.Println(add(3, 4))
	fmt.Println(swap("World", "Hello"))
	fmt.Println(split(1))
}
