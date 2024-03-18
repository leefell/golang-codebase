package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello!")
}

func variaveis() {
	// : é usado apenas na inferencia
	a := "Alexandre"
	b := 10
	c := 3.1233
	d := false

	fmt.Print("\n", a)
	fmt.Print("\n", b)
	fmt.Print("\n", c)
	fmt.Print("\n", d)

	fmt.Printf("\n %T", b)
}

func Soma(a int, b int) int {
	return a + b
}

func main() {
	resultado := Soma(1, 1)
	fmt.Printf("%T", resultado)
	fmt.Print("\n", resultado)
}
