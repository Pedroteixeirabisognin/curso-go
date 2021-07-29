package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	//Não é uma data, é a ordem que será passada dia,mes,ano, hora, minuto e segundo
	s := time.Now().Format("02/01/2006 03:04:05")

	fmt.Fprintf(w, "<h1>Hora certa: %s<h1>", s)

}

func main() {
	//Mapeamento da url
	//horaCerta é um modelo para uma função usada dentro de HandleFunc
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
