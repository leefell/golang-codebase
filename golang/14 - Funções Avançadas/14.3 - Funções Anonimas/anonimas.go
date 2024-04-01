package main

import "fmt"

func main() {

	func(texto string) {
		fmt.Println(texto)
	}("Função anônima")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebiido -> %s", texto)
	}("Passando parâmetro")

	fmt.Println(retorno)
}
