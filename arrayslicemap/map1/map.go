package main

import (
	"fmt"
)

func main() {
	//var aprovados map[int]string
	// map deve ser inicializado
	aprovados := make(map[int]string) // Cria um map com chave integer e valor string
	aprovados[12345678] = "Maria"
	aprovados[12355555] = "Pedro"
	aprovados[33332221] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[33332221])
	delete(aprovados, 33332221)
	fmt.Println(aprovados[33332221])
}
