package main

import (
	"fmt"
	"sync"
	"time"
)

func raceT() {
	r := 0
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(rid int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				time.Sleep(1 * time.Millisecond)
				fmt.Printf("%d|%d: %d\n", rid, i, r)
				r++
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Result: %d\n", r)
}
