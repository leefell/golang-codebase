package main

import (
	"errors"
	"fmt"
)

func main() {
	valorZero()
}

func numeros() {
	var numero int64 = -100000
	fmt.Println(numero)

	var numero2 uint32 = 1000000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1323.45
	fmt.Println(numeroReal2)

	numeroReal3 := 323.34
	fmt.Println(numeroReal3)
}

func strings() {
	var tipoString string = "string declara como string"
	fmt.Println(tipoString)

	// golang nao tem char
	char := 'A'
	fmt.Println(char) // nesse caso o compilador retorna o valor de 'A' da tabela ASCII
}

func valorZero() {
	// No go, todo tipo de dado tem seu valor zero
	var texto string
	fmt.Println(texto)

	var exemplo int32
	fmt.Println(exemplo)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
