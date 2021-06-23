package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y)) // É necessário converter para as variáveis ficarem iguáis, caso contrário dará erro

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado...
	fmt.Println("Teste " + string(123))

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123") // Retorna dois valores um número e um erro o _ seria para ignorar
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") // Qualquer tipo diferente de true irá retornar false
	if b {
		fmt.Println("Verdadeiro")
	}
}
