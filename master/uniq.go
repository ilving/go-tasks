package uniq

import (
	"fmt"
	"math/rand"
	"sync"
)

func generator() {
	var exist map[int]bool
	capacity := 1000

	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10))
	}

	uniqueIDs := make(chan int)
	wg := sync.WaitGroup{}

	for i := 0; i < capacity; i++ {
		i := i

		wg.Add(1)
		go func() {
			if _, ok := exist[i]; !ok {
				exist[doubles[i]] = true

				uniqueIDs <- doubles[i]
			}
			wg.Done()
		}()
		wg.Wait()

		for val := range uniqueIDs {
			fmt.Println(val)
		}
	}
}
