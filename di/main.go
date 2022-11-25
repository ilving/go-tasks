package main

import (
	bottom "github.com/ilving/test/di/bottom"
	top "github.com/ilving/test/di/top"
)

func main() {
	l2 := bottom.NewL2()

	l1 := top.NewLayer1(l2)

	print(l1)
}
