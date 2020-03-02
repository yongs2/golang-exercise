package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := os.Getenv("MARIADB_USER")
	dbPwd := os.Getenv("MARIADB_PWD")
	dbAddr := os.Getenv("MARIADB_HOST") + ":" + os.Getenv("MARIADB_PORT")
	dbName := os.Getenv("MARIADB_NAME")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPwd, dbAddr, dbName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO names VALUES (?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(nil, "John")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT id, name FROM names")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var id int
	var name string
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Printf("%d: %s\n", id, name)
	}
}
