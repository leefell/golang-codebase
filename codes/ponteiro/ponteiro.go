package main

import "fmt"

type Carro struct {
	Name string
}

// Para alterar valores que estão rodando nessa instancia voce coloca o *

func main() {
	carro := Carro{
		Name: "Honda Civic",
	}

	carro.andou()
	fmt.Println(carro.Name)
}

func (c *Carro) andou() {
	c.Name = "Fiesta"
	fmt.Println(c.Name, "andou!")
}

// -------------------------------------------------------------------------

func notes() {
	a := 10
	fmt.Println("Endereco da variavel A: ", &a)

	var ponteiro *int = &a
	fmt.Println("Endereco sendo apontado pelo 'ponteiro': ", ponteiro)
	fmt.Println("Valor do endereço que esta sendo apontado por 'ponteiro': ", *ponteiro)

	// Como estou alterando o valor do endereco de memoria pelo ponteiro, a variavel 'a' também sofre alteração;
	*ponteiro = 20
	fmt.Println("Novo valor atribuido diretamente no ponteiro: ", *ponteiro)
}

func valorSerAlterado() {
	variavel := 10
	fmt.Println("Antes da chamada da função:", variavel)

	alterarValor(&variavel)
	fmt.Println("Depois da chamada da função:", variavel)
}

func alterarValor(a *int) {
	*a = 1000
}
