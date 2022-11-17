package main

func visT() {
	a := 1

	// {
	// 	a := 2
	// 	print(a)
	// }

	// print(a)
	defer print(a)

	defer func() {
		print(a)
	}()

	a = 2
}
