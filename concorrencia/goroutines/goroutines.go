package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// Código síncrono
	// fale("Maria", "Pq vc não fala comigo?", 3)
	// fale("João", "Só posso falar depois de vc!", 1)

	// Código assincrono, porém ele será chamado e passará para o próximo método, se não houver
	// as goroutines são encerradas antes de executarem e encerra o programa
	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)
	// time.Sleep(time.Second * 5)
	// fmt.Println("Fim!!!")

	go fale("Maria", "Entendi", 10)
	fale("João", "Parabéns!", 5)

}
