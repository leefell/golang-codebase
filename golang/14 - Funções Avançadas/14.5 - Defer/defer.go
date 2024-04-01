package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovavdo(n1, n2 float32) bool {
	// essa linha sera executada antes de qualquer um dos returns
	defer fmt.Println("Média calculada. Resultado será retornado")

	fmt.Println("Entrando na função que verifica se um aluno está aprovado.")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

// Defer adia a execução de uma função até o último momento possível
func main() {

	if alunoEstaAprovavdo(8, 8) {
		fmt.Println("O aluno foi aprovado!")
	} else {
		fmt.Println("Aluno reprovado!")
	}
}
