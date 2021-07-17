package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	// * é usado para desreferenciar, ou seja ter acesso ao valor no qual o ponteiro aponta
	*n++
}
func main() {
	n := 1
	inc1(n) // Por valor
	fmt.Println(n)

	// & é usado para obter o endereço da variável
	inc2(&n) //Por referência
	fmt.Println(n)
}
