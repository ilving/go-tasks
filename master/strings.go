package main

import "fmt"

func strF() {
	// -------------------

	//s := "123"
	s := "asd漢字"
	for k, v := range s {
		fmt.Println(k, v)
	}

	s := "asd漢字"
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	// -------------------

	var strData string
	for i := 0; i < 10; i++ {
		strData += fmt.Sprintf("%d", i)
	}
	fmt.Println(strData)

	bs := []byte{97, 115, 100, 230, 188, 162, 229, 173, 151} // "asd漢字"
	println(string(bs))
	res := []rune(string(bs))
	for i := range res {
		if res[i] >= 0x7F {
			res[i] = '?'
		}
	}

	println(string(res))

	res = make([]rune, 0, len(bs))
	for i := 0; i < len(bs); i++ {
		if bs[i] < 127 {
			res = append(res, rune(bs[i]))
		} else {
			res = append(res, '?')
			switch {
			case bs[i]>>4 == 15:
				i += 3
			case bs[i]>>5 == 7:
				i += 2
			case bs[i]>>6 == 3:
				i += 1
			}
		}
	}
	println(string(res))
}
