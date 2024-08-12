package main

import "time"

func main() {
	sl := make([]bool, 1000)

	for i := range sl {
		go func(i int) {
			sl[i] = true
		}(i)
	}

	time.Sleep(1 * time.Second)
}
