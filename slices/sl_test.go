package main

import (
	"testing"
	"unsafe"
)

const len = 100_000_000

func Benchmark_Append(b *testing.B) {
	var sl []int
	for i := 0; i < b.N; i++ {
		sl = make([]int, 0, len)
		for k := 0; k < cap(sl); k++ {
			sl = append(sl, k)
		}
	}
}

func Benchmark_Idx(b *testing.B) {
	var sl []int
	for i := 0; i < b.N; i++ {
		sl = make([]int, len, len)
		for k := 0; k < cap(sl); k++ {
			sl[k] = k
		}
	}
}

func Benchmark_Unsafe(b *testing.B) {
	var sl []int
	var i int
	eleSize := int(unsafe.Sizeof(i))

	for i := 0; i < b.N; i++ {
		sl = make([]int, len, len)
		p := unsafe.Pointer(&sl[0])
		for k := 0; k < cap(sl); k++ {
			*(*int)(unsafe.Add(p, k*eleSize)) = k
		}
	}
}
