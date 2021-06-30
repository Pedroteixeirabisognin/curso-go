package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) { //Sempre vai retornar na ordem certinha do nome das variáveis, serve para organizar o retorno
	segundo = p2
	primeiro = p1

	return // retorno limpo
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
