package main

import "fmt"

func main() {
	doSomething(4)
}

func doSomething(n int) {
	defer fmt.Println(n)
	fmt.Println(n)
}
