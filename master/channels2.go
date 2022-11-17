package main

import (
	"fmt"
)

func channels2() {
	ch := make(chan int, 10)

	for i := 0; i < cap(ch)/2; i++ {
		ch <- i
	}

	for i := 0; i < 2*cap(ch); i++ {
		v, f := <-ch
		fmt.Println(v, f)
	}
}
