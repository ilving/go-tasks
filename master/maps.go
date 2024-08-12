package main

import (
	"fmt"
)

func mapsF(v map[int]int) {
	v[1] = 1
}

func main() {
	m0 := map[int]int{0: 0, 2: 2}

	mapsF(m0)
	fmt.Printf("%+v\n", m0)
}

func mapsT2() {
	// x32
	type myStruct struct {
		b1 byte
		b2 byte
		v  int
	}

	m := map[int]struct{ v int }{0: {v: 1}}
	m[0].v = 0
	print(m[0].v)

	s := []myStruct{{v: 1}}
	s[0].v = 0 // slice.pointer + index*sizeof(struct) + offset(v)

	ms := map[int][]byte{}
	ms[0] = []byte{1, 2}
	ms[0][1] = 2 // {{len: .. cap: ... ptr: *unsafe}}
}
