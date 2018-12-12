package main

import "fmt"

func main() {
	type text string
	var t text = "gopher is me"
	fmt.Printf("type:%T, value: %q\n", t, t)
}
