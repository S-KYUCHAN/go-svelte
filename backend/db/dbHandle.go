package db

import (
	"fmt"
	"log"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbQuery(query string) *sql.Rows {
	dataSource := "kyu:admin@tcp(localhost:3306)/testdb"
	conn, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	rows, err := conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Endpoint Hit: DbQuery")
	return rows
}

func DbExec(query string) sql.Result {
	dataSource := "kyu:admin@tcp(localhost:3306)/testdb"
	conn, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	result, err := conn.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Endpoint Hit: DbQuery")
	return result
}


func DbQueryContext(query string) *sql.Rows {
	dataSource := "kyu:admin@tcp(localhost:3306)/testdb"
	conn, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	rows, err := conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Endpoint Hit: DbQueryContext")
	return rows
}