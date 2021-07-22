package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				//Cada looping ele vai parar nesse channel
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	//Quando ele terminar de executar o for ele vai para o close
	//É necessário encerrar os channels para não dar deadlock
	close(c)
}

func main() {

	c := make(chan int, 30)

	//Cap(c) indica a capacidade do channel que no caso foi setado 30
	//Deixa preparado 30 channels no for para serem executados
	go primos(cap(c), c)

	//Pega channel por channel
	//primos() vai ser executado aos poucos e uma única vez
	//Cada passada do looping ele vai executar o for do primos() até voltar no c <- primo
	for primo := range c {
		fmt.Printf("%d ", primo)
	}

	fmt.Println("Fim!")
}
