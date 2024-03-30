package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa    // essa seria a esquematização de uma herança
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Alexandre", "Augusto", 19, 1.78}
	fmt.Println(p1)

	e1 := estudante{p1, "BSI", "IFSP"}
	fmt.Println(e1)
}
