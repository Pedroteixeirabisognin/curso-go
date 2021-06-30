package main

import "fmt"

func closure() func() {
	x := 10
	var funcao = func() {
		// O go usa closure, ele sabe o contexto mais relacionado a função que possui
		//aquela variável chamada, em algumas outras linguagens daria erro dizendo que x não está declarado, por culpa do escopo diferente da função
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX()
}
