package main

import "fmt"

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	//Como está no mesmo pacote não existe visibilidade privada
	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}
