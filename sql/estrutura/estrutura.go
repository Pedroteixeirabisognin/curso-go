package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}

	return result
}

func main() {

	//Driver, usuario e senha com "@/" no final
	db, err := sql.Open("mysql", "root:MySql2019!@/")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	exec(db, "create database if not exists testedb")
	exec(db, "use testedb")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios(
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (ID)
	)`)
}
