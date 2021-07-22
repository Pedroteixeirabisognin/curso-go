package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines
// channel é um tipo

func doisTreisQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base

}

func main() {
	c := make(chan int)
	go doisTreisQuatroVezes(2, c)

	//Espera o channel ficar pronto pra continuar executando as outras linhas
	//Assim você consegue fazer esperar a goroutine
	a, b := <-c, <-c

	fmt.Println(a, b)
	fmt.Println(<-c)

	//Gera um deadlock por que não tem mais goroutines sendo executadas depois da linha anterior
	//Tente testar sem adicionar mais um c <- 4 * base em doisTresQuatroVezes
	// e depois tente adicionando note que o channel será executado baseado na quantidade
	// de goroutines existentes
	//fmt.Println(<-c)
}
