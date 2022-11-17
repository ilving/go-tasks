package main

import "fmt"

func slicesF(v []byte) {
	v[0] = 0
	v = append(v, 3)
}

func slicesT() {
	s0 := []byte{1, 2}
	slicesF(s0)
	fmt.Printf("%+v\n", s0)

	// ------------
	s1 := make([]byte, 2, 4)
	s1[0], s1[1] = 1, 2
	slicesF(s1)
	fmt.Printf("%+v\n", s1)
	// ------------

	var s2 []byte
	slicesF(s2)
	fmt.Printf("%+v\n", s2)
}
