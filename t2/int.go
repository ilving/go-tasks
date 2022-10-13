package main

import (
	"fmt"
	"time"
)

func intT() {

	var i int8
	for {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(i)
		i++
	}
}
