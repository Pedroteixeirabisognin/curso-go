package main

import "fmt"

type curso struct {
	nome string
}

func main() {

	//pode ser usado em conjunto do switch para verificar o tipo do valor interno
	//e fazer tarefas diferentes baseadas nisso
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a Linguagem do Google"}
	fmt.Println(coisa2)
}
