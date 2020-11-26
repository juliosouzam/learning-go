package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@/curso_go")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("insert into users(name) values(?)")
	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")

	id, _ := res.LastInsertId()

	fmt.Println(id)

	linhas, err := res.RowsAffected()
	fmt.Println(linhas)
}
