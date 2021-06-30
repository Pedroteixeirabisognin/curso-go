package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // tipo, declarados, capacidade
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s)

	s = append(s, 1) // Slice adiciona referências automaticamente se passar do lenght e dobra a capacidade
	fmt.Println(s, len(s), cap(s))

	s = append(s, 2) // Len = 22, Cap = 40
	fmt.Println(s, len(s), cap(s))

}
