// Programas executáveis iniciam pelo pacote main
package main

/*
Os códigos em Go são organizados em pacotes para usá-los é necessário declarar um ou vários imports
O auto formatador vai sumir com imports que não estão sendo utilizados
*/

import "fmt"

func main() {
	fmt.Print("Primeiro")
	fmt.Print("Programa!")
}
