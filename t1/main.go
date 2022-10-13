package main

import (
	"math"
	"sort"
	"time"
)

const lim = 1000000

func sSort(list []int) int {
	sort.Ints(list)
	return list[len(list)-1] + list[len(list)-2]
}

func sList(list []int, amountMax int) int {
	m := make([]int, amountMax)
	for i := range m {
		m[i] = -1 * math.MaxInt
	}

	for _, v := range list {
		if m[0] < v {
			for j := 1; j < amountMax; j++ {
				m[j] = m[j-1]
			}
			m[0] = v
		}
	}
	return m[0] + m[1]
}

// func main() {
// 	ls := make([]int, lim)
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < lim; i++ {
// 		ls[i] = rand.Int()
// 	}

// 	//ls = []int{1, 2, 19, 5, 4, 11}
// 	var st int64
// 	var r int

// 	st = time.Now().UnixMicro()
// 	r = sSort(ls)
// 	st = time.Now().UnixMicro() - st
// 	fmt.Printf("sSort: %d %fs\n", r, float64(st)/1000000.0)

// 	st = time.Now().UnixMicro()
// 	r = sList(ls, 2)
// 	st = time.Now().UnixMicro() - st
// 	fmt.Printf("sList: %d %fs\n", r, float64(st)/1000000.0)
// }

func main() {
	defer func() {
		if err := recover(); err != nil {
			print(err)
		}
	}()

	go func() {
		panic("test")
	}()

	time.Sleep(10 * time.Second)

	print("test2")
}
