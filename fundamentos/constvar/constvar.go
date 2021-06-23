package main

import (
	"fmt"
	//Referência reduzida. Se usar _ pode importar sem usar
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é ", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// true vai pra variável "e" e false vai pra variável "e"
	var e, f bool = true, false
	fmt.Println(e, f)

	//Usando a forma reduzida de criar var fica como se fosse g = 2,  h = false e i = "epa!"
	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
