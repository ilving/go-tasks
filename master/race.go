package main

import (
	"fmt"
)

func raceT() {
	r := 0

	for i := 0; i < 3; i++ {
		go func() { r += i }()
	}

	fmt.Printf("Result: %d\n", r)
}
