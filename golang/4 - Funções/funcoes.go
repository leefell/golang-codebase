package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadoSoma, resultadosSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadosSubtracao)

	// Quando desejamos extrair apenas um resultado específico de uma função que retorna múltiplos valores,
	// podemos utilizar o '_' para ignorar os valores que não necessitamos.
	resultadoSoma2, _ := calculosMatematicos(10, 30)
	fmt.Println(resultadoSoma2)
}
