package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 10 {
		i++
		//fmt.Println("Incrementado i: ", i)
		//time.Sleep(time.Second)
	}

	fmt.Println("O último valor de 'i' é: ", i)

	for j := 0; j < 10; j++ {
		//fmt.Println("Incrementando j: ", j)
		//time.Sleep(time.Second)
	}

	nomes := [3]string{"Alexandre", "Marcelo", "Gabriel"}

	// Por padrão o primeiro parametro é sempre o indice, o segundo é o valor
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Alexandre",
		"sobrenome": "Augusto",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// Para executar um loop infinito é só colocar for{}
}
