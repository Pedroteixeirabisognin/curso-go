package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//Driver, usuario e senha com "@/" no final
	db, err := sql.Open("mysql", "root:MySql2019!@/testedb")
	if err != nil {
		panic(err)
	}

	//Fecha o banco ao final da função
	defer db.Close()

	stmt, _ := db.Prepare("insert into usuarios(nome) values (?)")
	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)
}
