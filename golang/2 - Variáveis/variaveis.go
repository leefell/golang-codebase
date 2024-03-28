package main

import "fmt"

func main() {
	var variavel1 string = "Texto" // explicito
	variavel2 := "Texto 2"         // inferindo

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "blablabla"
		variavel4 string = "blablabla"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "var5", "var6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5 // inverteu os valores das variaveis sem usar um aux
	fmt.Println(variavel5, variavel6)
}
