package main

import (
	"fmt"
)

func intT() {

	for i := int8(0); i < i+1; i++ {
		fmt.Println(i)
	}

	for i := 0; ; i++ {
		if i > i+1 {
			fmt.Println(i)
			break
		}
	}

}
