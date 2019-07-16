package main

import "fmt"

func main() {
	up, get := factory()
	up()
	up()
	fmt.Println(get())

	up()
	fmt.Println(get())
}

func factory() (func(), func() int) {
	var i int
	return func() {
			i++
		},
		func() int {
			return i
		}
} // END
