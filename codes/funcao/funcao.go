package main

import "fmt"

// Definição da estrutura Carro que possui um campo Nome
type Carro struct {
	Nome string
}

// Método andar associado à estrutura Carro, que imprime uma mensagem indicando que o carro andou
func (c Carro) andar() {
	fmt.Println(c.Nome, "andou")
}

func main() {
	// Criação de uma instância da estrutura Carro com o nome "Honda Civic"
	carro := Carro{
		Nome: "Honda Civic",
	}

	carro.andar()

	resultado1, str1, str2 := soma(10, 20)
	fmt.Println(resultado1, str1, str2)

	resultado2 := subtracao(10, 5)
	fmt.Println(resultado2)

	resultado3 := somaArray(4, 2, 3, 54, 7, 5, 34, 2, 3, 4, 5, 6, 7, 8, 34)
	fmt.Println(resultado3)

	// Declaração e chamada de uma função anônima que retorna outra função
	resultado4 := func(x ...int) func() int {
		res := 0

		// Soma todos os valores passados como argumento
		for _, v := range x {
			res += v
		}

		// Retorna uma função anônima que retorna o quadrado da soma
		return func() int {
			return res * res
		}
	}

	// Chama a função anônima que retorna outra função e em seguida a função retornada
	fmt.Println(resultado4(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)())
}

// Funções podem retornar mais de um valor
func soma(a int, b int) (int, string, string) {
	return a + b, "somou! |", "outro valor sendo retornado"
}

// Return nomeado
func subtracao(a int, b int) (result int) {
	result = a - b
	return
}

// Função variádica x...int permite passar uma quantidade variável de argumentos inteiros
func somaArray(x ...int) int {
	resultado := 0

	for _, value := range x {
		resultado += value
	}

	return resultado
}
