package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string)

	for i := 0; i < rand.Intn(100000); i++ {
		wg.Add(1)
		go func(group sync.WaitGroup) {
			defer group.Done()

			ch <- fmt.Sprintf("Goroutine %d", i)
		}(wg)
	}

	wg.Wait()

	for v := range ch {
		fmt.Println(v)
	}

}
