package main

import (
	"fmt"
	"sync"
	"time"
)

// goRoutine são funções ou métodos que voce pode chamar a execucao e
// nao necessariamente esperar que elas terminem para chamar outra

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Olá mundo!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Aprendendo sobre waitGroup!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
