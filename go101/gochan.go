package main

import (
	"fmt"
	"time"
)

func main() {
	total := 10
	ch := make(chan struct{})
	now := time.Now()
	for i := 0; i < total; i++ {
		go printout(i, ch)
	}
	for i := 0; i < total; i++ {
		<-ch
	}
	fmt.Println(time.Now().Sub(now))
}

func printout(i int, ch chan struct{}) {
	fmt.Println(i)
	ch <- struct{}{}
}
