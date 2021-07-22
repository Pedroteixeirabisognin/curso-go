package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Executou!") //Ele é executado antes dos channels,
	//as vezes ele pode ser executado no meio por que chamou um channel só
	ch <- 4
	ch <- 5
	ch <- 6
}

func main() {
	// Ele vai ler 3 channels e esperar e assim por diante até acabar
	// e volta a adicionar a medida que vai executando as tarefas
	ch := make(chan int, 3)

	go rotina(ch)

	time.Sleep(time.Second)

	//Ele nunca vai executar o quinto por que o channel só é chamado 4 vezes
	//A cada chamada dessas ele vai executar um e puxar 3 que vão ficar executando em background

	//Você não ler todas as informações do channel
	//não tem problema ele só vai terminar o programa
	fmt.Println(<-ch)

	//Tente comentar e descomentar os codigos abaixo para testar

	//Vai executar 2 do buffer e puxar o 4 para o buffer ([3,4,5] no buffer)
	fmt.Println(<-ch)
	//Vai executar 3 do buffer e puxar 5 para o buffer ([4,5,6] no buffer)
	fmt.Println(<-ch)

	//Vai executar 4 do buffer ([5,6] no buffer)
	//As vezes executou pode aparecer aqui devido ao tempo de encerramento do programa
	fmt.Println(<-ch)
	//Vai executar "Executou!" por que passou do 4 e vai executar 5 ([6] no buffer)
	fmt.Println(<-ch)

	//Ele vai encerrar o programa sem executar o 6 do buffer por que não foi chamado um
	//novo channel

}
