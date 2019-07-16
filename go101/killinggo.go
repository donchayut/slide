package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	now := time.Now()
	go lazy(time.Second, "wait 1 second")
	go lazy(time.Second*2, "wait 2 seconds")
	lazy(time.Second*3, "wait 3 seconds")
	fmt.Println("usage time", time.Now().Sub(now))
}

func lazy(d time.Duration, s string) {
	time.Sleep(d)
	fmt.Println(s)
}
