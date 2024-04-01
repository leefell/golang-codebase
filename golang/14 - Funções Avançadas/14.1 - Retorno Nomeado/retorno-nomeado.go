package main

import "fmt"

func calculos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2

	// Com retorno nomeado voce pode dar só um return que o compilador ja sabe
	// o que deve ser retornado
	return
}

func main() {
	soma, subtracao := calculos(10, 10)
	fmt.Println(soma, subtracao)
}
