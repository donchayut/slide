package main

import "fmt"

func main() {
	show(area)
}

type areaFunc func(int, int) int

func area(dx, dy int) int {
	return dx * dy
}

func show(fn areaFunc) {
	fmt.Printf("area of 64x48 = %d", fn(64, 48))
}
