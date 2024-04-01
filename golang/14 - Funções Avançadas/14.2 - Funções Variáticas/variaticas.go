package main

import "fmt"

// Essa funcao recebe de 0 - N números
func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 2, 4, 4, 3, 2, 4, 6, 8)
	fmt.Println(totalSoma)

	escrever("Olha o texto!", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
