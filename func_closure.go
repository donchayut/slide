package main

import "fmt"

func main() {
	counter := counterFactory()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func counterFactory() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
