package main

import (
	"fmt"
	"strconv"
)

type rectangle struct {
	width  int
	length int
}

func (rec rectangle) area() int {
	return rec.width * rec.length
}

func (rec rectangle) String() string {
	return "Rectangle"
}

type color struct {
	r, g, b int
}

func (c color) YCML() (int, int, int, int) {
	return 0, 0, 0, 0
}

type rectangleColor struct {
	areaer
	colorer
}

type colorer interface {
	YCML() (int, int, int, int)
}

type areaer interface {
	area() int
}

func printArea(rec areaer) {
	fmt.Printf("%q\n", rec)
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) area() int {
	return int(0.5 * t.base * t.height)
}

type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func main() {
	runtime.GOMAXPROCS(1)
	// rec := rectangle{3, 4}
	rec := new(rectangle)
	rec.width = 3
	rec.length = 4
	printArea(rec)
	t := triangle{15, 4}
	printArea(t)

	var d Int = 999
	fmt.Printf("%q\n", d)
}
