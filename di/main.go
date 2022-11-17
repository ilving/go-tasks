package main

import (
	bottom "github.com/ilving/test/t3/bottom"
	top "github.com/ilving/test/t3/top"
)

func main() {
	l2 := bottom.NewL2()

	l1 := top.NewLayer1(l2)

	print(l1)
}
