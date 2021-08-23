package models

import (
	"database/sql"

	// "github.com/S-KYUCHAN/backend/modules/db"
)

type Base struct {
	TableName 	string
	Conn		*sql.DB
	Tx			*sql.Tx
}

func (b Base) SetConn(conn *sql.DB) Base {
	b.Conn = conn
	
	return b
}

// func (b Base) Table(table string) *db.SQL {
// 	return db.Table(table).WithDriver(b.Conn)
// }