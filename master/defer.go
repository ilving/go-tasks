package main

func deferT() {
	a := 1

	defer print(a)

	defer func() {
		print(a)
	}()

	a = 2
}
