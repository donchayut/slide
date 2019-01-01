package main

import "fmt"

func main() {
	doSomething(4)
}

func doSomething(n int) {
	defer fmt.Println(n)
	defer func() {
		fmt.Println(n)
	}()
	n = n * n
	fmt.Println(n)
}
