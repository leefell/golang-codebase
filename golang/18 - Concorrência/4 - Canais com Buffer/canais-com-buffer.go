package main

import "fmt"

func main() {
	// A diferenca do canal com buffer é que ele só vai bloquear quando atingir sua capacidade máxima
	canal := make(chan string, 3)
	canal <- "Agora são 12:15"
	canal <- "Mais uma string"
	canal <- "Terceiro valor"

	mensagem := <-canal
	mensagem2 := <-canal
	mensagem3 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)
}
