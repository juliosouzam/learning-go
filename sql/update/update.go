package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@/curso_go")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("update users set name = ? where id = ?")
	stmt.Exec("Geovana", 1)
	stmt.Exec("Roberta", 2)
}
