package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	//Enviando para o canal valor 1, valor inteiro
	ch <- 1 // Enviando dados para o canal (escrita)
	<-ch    //recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
