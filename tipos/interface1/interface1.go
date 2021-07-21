package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//interfaces são implementadas implicitamente

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

//Ele vai verificar se a struct que está entrando neste método possui
//o recever implementado na interface imprimível
func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	//Como pessoa possui um recever toString ele vai conseguir ser utilizado dentro de imprimir
	//Pois chama imprimivel que possui toString como um dos métodos declarados
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(produto{"Calça Jeans", 79.90})
}
