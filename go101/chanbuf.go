package main

import (
	"fmt"
)

func main() {
	total := 10
	ch := make(chan int, total)
	for i := total; i > 0; i-- {
		ch <- i
	}
	close(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
