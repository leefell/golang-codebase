package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {

		// Dessa maneira uma coisa não tem que esperar a outra, repare que o
		// time.sleep de um é diferente de outro, então
		// caso o canal 1 esteja pronto para receber um valor, printa ele na tela
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)

		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}
