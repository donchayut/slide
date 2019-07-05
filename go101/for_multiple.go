package main

import "fmt"

func main() {
	a := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(a)

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	fmt.Println(a)
}
