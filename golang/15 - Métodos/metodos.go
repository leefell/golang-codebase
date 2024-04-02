package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) salvar() {
	// Todos usuários terão um método salvar
	fmt.Printf("Salvando os dados do User %s no banco de dados!\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	user := usuario{"Fulano", 20}
	user.salvar()

	user2 := usuario{"Ciclano", 19}
	user2.salvar()

	fmt.Println(user.maiorDeIdade())
	fmt.Println(user2.maiorDeIdade())

	user2.fazerAniversario()
	fmt.Println(user2.idade)
}
