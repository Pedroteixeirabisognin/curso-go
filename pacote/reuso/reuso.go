package main

import (
	"fmt"

	"github.com/pedroteixeirabisognin/reuso/area"
)

func main() {
	fmt.Println(area.Circ(3.4))
}

//Para usar pacotes utilize crie um módulo com o comando
//go mod init github.com/pedroteixeirabisognin/nomeprojeto nomenclatura padrão
//criar módulos em src ficou deprecated desde a versão 1.13 do go.
