package main

import "fmt"

func main() {
	up, get := remoteCounterFactory()
	up()
	up()
	fmt.Println(get())

	up()
	fmt.Println(get())
}

func remoteCounterFactory() (func(), func() int) {
	var i int
	return func() {
			i++
		},
		func() int {
			return i
		}
}
