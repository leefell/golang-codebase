package main

// Canais são frequentemente utilizados para passar
// dados entre goroutines de forma segura.

// Até que um  canal receba um valor o programa não continua a rodar
import (
	"fmt"
	"time"
)

func main() {
	// Se o canal ainda estiver esperando dados e não recebe mais nenhum
	// acontece um erro chamado DeadLock, porque o programa fica eternamente
	// esperando um valor que não vai chegar.
	// O go não identifica um deadlock na compilação, apenas durante a execução
	// o que o torna perigoso.

	canal := make(chan string)
	go escrever("Testando canais", canal)

	fmt.Println("Depois da função escrever ser executada.")

	for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	// Sintaxe alternativa:
	// Enquanto o canal estiver aberto ele vai printando na tela
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa, canal fechado.")

}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 3)
	for i := 0; i < 5; i++ {
		canal <- texto // mandando um valor pra dentro do canal
		time.Sleep(time.Second)
	}

	close(canal)
}
