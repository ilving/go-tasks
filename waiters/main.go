package main

import (
	"fmt"
	"time"
)

func main() {
	go ready()
	go ready()
	go ready()

	Steady()

	time.Sleep(500 * time.Millisecond)
	//сделать что бы они начали бежать после вызова фунции, одновременно
}
func Steady() {
	// magic
}

func ready() {
	// ждем вызова стеди и запускаем все рутины одновременно

	fmt.Printf("%s\n", time.Now().Format(time.RFC3339Nano))
}
