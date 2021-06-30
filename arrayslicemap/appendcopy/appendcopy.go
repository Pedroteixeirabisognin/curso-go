package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6)
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2) // Cria um slice vazio de capacidade 2
	copy(slice2, slice1)     // copia dos dois primeiros valores do slice1 e insere no slice2
	fmt.Println(slice2)
}
