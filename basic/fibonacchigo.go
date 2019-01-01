package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	chNum := make(chan int)
	chQ := make(chan struct{})

	go fibonacchi(chNum, chQ)

	for v := range chNum {
		fmt.Println(v)
		if v > 1000 {
			close(chNum)
			chQ <- struct{}{}
		}
	}

	// chQ <- struct{}{}
}

func fibonacchi(chNum chan int, chQ chan struct{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	a, b := 0, 1

	// panic(errors.New("panic in go routine"))
	for {
		select {
		case chNum <- a:
			a, b = b, a+b
		case <-chQ:
			return
		}
	}
}
