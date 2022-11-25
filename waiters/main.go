package main

func runnerStart() {

	go ready()
	go ready()
	go ready()

	Steady()
	//сделать что бы они начали бежать после вызова фунции, одновременно
}
func Steady() {
	// magic
}

func ready() {
	// ждем вызова стеди

	// run()
}
