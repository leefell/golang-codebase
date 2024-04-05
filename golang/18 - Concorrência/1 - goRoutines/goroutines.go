package main

import (
	"fmt"
	"time"
)

func main() {
	// Concorrencia != paralelismo

	// goRoutine são funções ou métodos que voce pode chamar a execucao e
	//	nao necessariamente esperar que elas terminem para chamar outra
	go escrever("Olá mundo!")
	escrever("Aprendendo sobre concorrencia!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
