package main

import (
	"fmt"
)

func aaa(strings ...string) {
	strings = append(strings, "+a")
	fmt.Printf("%d", cap(strings))
	bbb(strings...)
}

func bbb(strings ...string) {
	strings = append(strings, "+b")
	fmt.Printf("%d %+v", cap(strings), strings)
}

func main() {

	slice := []string{"a", "b", "c"}
	slice = append(slice, "d")

	fmt.Printf("%d", cap(slice))

	aaa(slice...)
}
