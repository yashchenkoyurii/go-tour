package main

import (
	"fmt"
	"sync"
	"time"
)

func sum(n []int, c chan int) {
	var sum int

	for _, v := range n {
		sum += v
	}

	c <- sum
}

func fibonacci(c chan int, q chan int) {
	a, b := 0, 1

	for {
		select {
		case c <- a:
			a, b = b, a+b
		case <-q:
			return
		}
	}
}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Val(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()

	return c.v[key]
}

func main() {
	//Split calculation with goroutines
	n := []int{-5, 12, 15, -15, 14, 15}

	c := make(chan int, 2)

	go sum(n[:len(n)/2], c)
	go sum(n[len(n)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y)

	//Dead lock
	ch1 := make(chan int, 2)

	// go func(ch1 chan int) {
	// 	for range ch1 {
	// 		<-ch1
	// 	}
	// }(ch1)

	ch1 <- 1
	ch1 <- 2
	// ch1 <- 3

	//Range and close
	fibChan := make(chan int)

	go func(c chan int) {
		defer close(c)

		a, b := 0, 1

		for i := 0; i < 10; i++ {
			a, b = a+b, a

			c <- b
		}
	}(fibChan)

	for i := range fibChan {
		fmt.Println(i)
	}

	//Select
	fmt.Println("SELECT")
	ch3 := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch3)
		}

		quit <- 0
	}()
	fibonacci(ch3, quit)

	//MUTEX
	fmt.Println("Mutex")
	muxc := &SafeCounter{v: make(map[string]int)}

	for j := 0; j < 10000; j++ {
		go muxc.Inc("some key")
	}

	time.Sleep(time.Second)

	fmt.Println(muxc.Val("some key"))
}
