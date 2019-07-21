package main

import (
	"fmt"
)

func main() {
	fibonacci(10)

	fi := fiboclosure()
	for i := 0; i < 10; i++ {
		fmt.Println(fi())
	}

	chfi := make(chan int)
	chq := make(chan struct{})

	go fibogo(chfi, chq)
	for i := 0; i < 10; i++ {
		fmt.Println(<-chfi)
	}
	chq <- struct{}{}
}

func fibogo(fi chan int, q chan struct{}) {
	a, b := 0, 1

	for {
		select {
		case <-q:
			return
		case fi <- a:
			a, b = b, a+b
		}
	}
}

func fiboclosure() func() int {
	a, b := 0, 1
	return func() int {
		defer func() {
			a, b = b, a+b
		}()
		return a
	}
}

func fibonacci(n int) {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}
