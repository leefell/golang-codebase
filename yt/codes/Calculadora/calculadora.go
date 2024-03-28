package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	menu()

	reader := bufio.NewReader(os.Stdin)
	opcaoStr, _ := reader.ReadString('\n')

	opcao, err := strconv.Atoi(opcaoStr)
	if err != nil {
		fmt.Println("Erro ao ler a opção do menu: ", err)
		return
	}

	fmt.Println("Insira o primeiro número: ")
	numero1Str, _ := reader.ReadString('\n')
	numero1, err := strconv.Atoi(numero1Str)
	if err != nil {
		fmt.Println("Erro ao ler o primeiro número: ", err)
		return
	}

	fmt.Println("Insira o segundo número: ")
	numero2Str, _ := reader.ReadString('\n')
	numero2, err := strconv.Atoi(numero2Str)
	if err != nil {
		fmt.Println("Erro ao ler o segundo número: ", err)
		return
	}

	fmt.Println("Dados capturados com sucesso.")

	switch opcao {
	case 1:
		fmt.Println("Resultado da soma:", soma(numero1, numero2))
	case 2:
		fmt.Println("Resultado da subtração:", subtrair(numero1, numero2))
	case 3:
		fmt.Println("Resultado da divisão:", dividir(numero1, numero2))
	case 4:
		fmt.Println("Resultado da multiplicação:", multiplicar(numero1, numero2))
	case 5:
		fmt.Println("Saindo...")
	default:
		fmt.Println("Opção inválida.")
	}
}

func menu() {
	fmt.Println("Calculadora de dois números: ")
	fmt.Println("[1] - Somar ")
	fmt.Println("[2] - Subtrair ")
	fmt.Println("[3] - Dividir ")
	fmt.Println("[4] - Multiplicar ")
	fmt.Println("[5] - Sair")
}

func soma(a int, b int) int {
	return a + b
}

func subtrair(a int, b int) int {
	return a - b
}

func multiplicar(a int, b int) int {
	return a * b
}

func dividir(a int, b int) int {
	return a / b
}
