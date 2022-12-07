package main

func visT() {
	a := 1

	{
		a := 2
		print(a)
	}

	print(a)
}
