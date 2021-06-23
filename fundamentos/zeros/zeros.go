package main

import "fmt"

func main() {
	//É necessário definir o tipo quando não é recebido valor
	var a int     // 0
	var b float64 // 0
	var c bool    // false
	var d string  // ""
	var e *int    // <nil>

	fmt.Printf("%v %v %v %v %v", a, b, c, d, e)
}
