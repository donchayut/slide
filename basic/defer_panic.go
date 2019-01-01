package main

import "fmt"

func main() {
	catchMe()
}

func catchMe() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	s := []int{}

	fmt.Println(s[1])
}
