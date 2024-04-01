package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	//numero := 20
	//numeroInvertido := inverterSinal(numero)
	//fmt.Println(numeroInvertido)
	//fmt.Println(numero)

	fmt.Println("Alterando diretamente na variavel")

	novoNumero := 40
	fmt.Println("Antes da função:", novoNumero)

	inverterSinalComPonteiro(&novoNumero)
	fmt.Println("Após a função: ", novoNumero)
}
