package main

import (
	"log"
	"net/http"

	"github.com/serverdb/cliente"
)

func main() {
	http.HandleFunc("/usuarios/", cliente.UsuarioHandler)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
