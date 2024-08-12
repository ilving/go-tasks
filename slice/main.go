package main

import "fmt"

func main() {
	a := []int{}
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)

	b := append(a, 4)
	c := append(a, 5)

	fmt.Println(a, b, c)

	// ---------------------------------

	type myStruct struct {
		vp *int
		v  int
	}

	var vp *int
	vp = &a[0]
	m0 := map[interface{}]int{
		myStruct{v: 0, vp: vp}: 1,
		myStruct{v: 0}:         0,
	}

	//	mapsF(m0)
	fmt.Printf("%+v\n", m0)
}
