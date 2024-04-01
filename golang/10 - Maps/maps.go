package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Alexandre",
		"sobrenome": "Augusto",
	}

	fmt.Println("Nome: ", usuario["nome"])
	fmt.Println("Sobrenome: ", usuario["sobrenome"])
	fmt.Println(usuario)

	// Map com chave tipo string e o valor da chave é outro map
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Gabriel",
			"segundo":  "Tangerina",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Votuporanga",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["idade"] = map[string]string{
		"chaveIdade": "19",
	}

	fmt.Println(usuario2)

}
