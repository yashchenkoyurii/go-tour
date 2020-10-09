package main

import (
	"fmt"
	"io"
	"math"
	"strings"
)

type Vertex struct {
	X float64
	Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type myFloat float64

func (f myFloat) Abs() myFloat {
	if f < 0 {
		return -f
	}

	return f
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := myFloat(-32.5)
	fmt.Println(f.Abs())

	var i I
	i = &T{"Hello world!"}
	i.M()
	describe(i)

	var i1 interface{} = "Hello world"

	s, ok := i1.(string)
	fmt.Println(s, ok)

	f1, ok := i1.(float64)
	fmt.Println(f1, ok)

	switch val := i1.(type) {
	case string:
		fmt.Println("string: ", val)
	case interface{}:
		fmt.Println("interface: ", val)
	}

	r := strings.NewReader("Hello world!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("b[:%d] = %q\n", n, b[:n])
		if err == io.EOF {
			break
		}
	}
}
