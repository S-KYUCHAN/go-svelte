package models

import (
	// "fmt"
	// "strconv"
	"database/sql"
)

type UserModel struct {
	Base							`json:"-"`

	Id            int64             `json:"id"`
	Name          string            `json:"name"`
	UserName      string            `json:"user_name"`
	Password      string            `json:"password"`

	// CreatedAt 	  string 			`json:"created_at"`
	// UpdatedAt 	  string 			`json:"updated_at"`
};

func User() UserModel {
	return UserModel{Base: Base{TableName: "users"}}
}

func (u UserModel) SetConn(conn *sql.DB) UserModel {
	u.Conn = conn
	return u
}

func (u UserModel) IsEmpty() bool {
	return u.Id == int64(0)
}

// func (u UserModel) WithTx(tx *sql.tx) UserModel {
// 	u.Tx = tx
// 	return u
// }

// func (u UserModel) Find(id interface{}) UserModel {
// 	// item, _ := u.Table()
// }