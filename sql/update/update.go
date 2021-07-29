package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//Driver, usuario e senha com "@/" no final
	db, err := sql.Open("mysql", "root:MySql2019!@/testedb")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//Não passe valor passado pelo usuário concatenando string, cuidado com sql injection
	//Update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Uóxiton Istive", 1)
	stmt.Exec("Sherinstone Uasleska", 2)

	//delete
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)
}
