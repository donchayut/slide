package main

import "fmt"

func main() {
	fi := fibonacchiFunc()
	for i := 0; i < 10; i++ {
		fmt.Println(fi())
	}
}

func fibonacchiFunc() func() int {
	a, b := 0, 1
	return func() int {
		defer func() {
			a, b = b, a+b
		}()
		return a
	}
}
