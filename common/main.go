package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(1)
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(2)
		}
	}()

	time.Sleep(5 * time.Second)
}
