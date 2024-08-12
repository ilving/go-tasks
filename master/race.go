package main

import (
	"fmt"
	"time"
)

func main() {
	r := 0

	for i := 0; i < 3; i++ {
		go func() { r += i }()
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("Result: %d\n", r)
}
