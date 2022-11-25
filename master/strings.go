package main

import "fmt"

func strF() {
	// -------------------

	//s := "123"
	s := "W漢字"
	for k, v := range s {
		fmt.Println(k, v)
	}

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	// -------------------

	var strData string = ""
	for i := 0; i < 1000; i++ {
		strData += fmt.Sprintf("%d", i)
	}
	fmt.Println(strData)
}
