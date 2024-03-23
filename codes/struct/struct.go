package main

import(
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct{
	Nome  string
	Email string
	CPF   int 
}

// Composição de uma struct com outra
type ClienteInternacional struct{
	Cliente 
	Pais  string `json:"pais"` // quando rodar a funcao, ele transforma o Pais em pais
}

// Atachando uma função em um struct, se tornando um método
func (c Cliente)Exibe(){
	fmt.Println("Exibindo cliente pelo metodo: ", c.Nome)
}

func main(){
	
	cliente := Cliente{
		Nome: "Alexandre",
		Email: "alexandre@gmail.com",
		CPF: 12345678990,
	}
	fmt.Println(cliente)
	
	cliente2 := Cliente{"Andre", "andre@gmail.com", 1242312353}
	fmt.Printf("Nome: %s. Email: %s. CPF: %d\n", cliente2.Nome, cliente2.Email, cliente2.CPF)

	cliente3 := ClienteInternacional{
		Cliente: Cliente{
			Nome: "Gabriel",
			Email: "tang@hotmail.com",
			CPF: 12345612990,
		},
		Pais: "Brasil",
	}
	fmt.Printf("Nome: %s. Email: %s. CPF: %d Pais: %s\n", cliente3.Nome, cliente3.Email, cliente3.CPF, cliente3.Pais)

	// Parece muito com uma herança é uma composição
	cliente.Exibe()
	cliente2.Exibe()
	cliente3.Exibe()

	// Marshal pega os dados em um slice de bytes
	clienteJson, err := json.Marshal(cliente3)
	if err != nil{
		log.Fatal(err.Error())
	}

	fmt.Println((string(clienteJson)))

	// "Desempacotando" esse json
	jsonCliente4 := `{"Nome":"Gabriel","Email":"tang@hotmail.com","CPF":12345612990,"pais":"Brasil"}`
	cliente4 := ClienteInternacional{}

	json.Unmarshal([]byte(jsonCliente4), &cliente4)
	fmt.Println(cliente4)
}