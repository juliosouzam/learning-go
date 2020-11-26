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

	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("insert into users(id,name) values(?,?)")
	stmt.Exec(3000, "Bia")
	stmt.Exec(3001, "Carlos")
	_, er := stmt.Exec(1, "Thiago")

	if er != nil {
		tx.Rollback()
		log.Fatal(er)
	}

	tx.Commit()
}
