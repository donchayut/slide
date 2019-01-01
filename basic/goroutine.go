package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(100)
	now := time.Now()
	for i := 0; i < 100; i++ {
		go justPrint(i)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(now))
}

func justPrint(v interface{}) {
	time.Sleep(time.Millisecond * 100)
	fmt.Println(v)
	wg.Done()
}
