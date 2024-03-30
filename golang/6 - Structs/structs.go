package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "Alexandre"
	u.idade = 19
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua Numero tal", 0}

	usuario2 := usuario{"Gabriel", 19, enderecoExemplo }
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 22}
	fmt.Println(usuario3)
}
