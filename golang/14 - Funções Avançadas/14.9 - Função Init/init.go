package main

import "fmt"

var n int

func init() {
	// essa função sempre vai ser a primeira a ser executada
	// limitado à uma por arquivo
	fmt.Println("Função Init sendo executada!")
	n = 10
}

func main() {
	fmt.Println("Função Main em execução!")
	fmt.Println(n)
}
