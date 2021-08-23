package db

import (
	"fmt"
	"log"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	// "github.com/S-KYUCHAN/backend/models"
)

func InitDB() *sql.DB {
	dataSource := "kyu:admin@tcp(localhost:3306)/testdb"
	conn, err := sql.Open("mysql", dataSource)
	if err != nil {
		if conn != nil {
			_ = conn.Close()
		}
		panic(err)
	}

	if err := conn.Ping(); err != nil {
		panic(err)
	}
	
	return conn
}

// func DbQuery(conn *sql.DB, query string) *sql.Rows {
func DbQuery(query string) *sql.Rows {

	conn := InitDB()
	defer conn.Close()

	rows, err := conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Endpoint Hit: DbQuery")
	return rows
}

func DbQueryRow(query string, user ...interface{}) {
	
	conn := InitDB()
	defer conn.Close()

	err := conn.QueryRow(query).Scan(user... )
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Endpoint Hit: DbQuery")
}

// func DbExec(conn *sql.DB, query string) sql.Result {
func DbExec(query string) sql.Result {

	conn := InitDB()
	defer conn.Close()

	result, err := conn.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Endpoint Hit: DbQuery")
	return result
}


func DbQueryContext(query string) *sql.Rows {

	conn := InitDB()
	defer conn.Close()

	rows, err := conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Endpoint Hit: DbQueryContext")
	return rows
}