package models

import (
	"database/sql"

	// "github.com/S-KYUCHAN/backend/modules/db"
)

type ArticleModel struct {
	Base 						`json:"-"`
	Id				int64		`json:"id"`
	Title			string		`json:"title"`
	Description		string		`json:"description"`
	Content			string		`json:"content"`

	CreatedAt 		string 		`json:"created_at"`
	UpdatedAt 		string 		`json:"updated_at"`
}

func Article() ArticleModel {
	return ArticleModel{Base: Base{TableName: "articles"}}
}

func (a ArticleModel) GetAriticle() ArticleModel {
	return a
}

func (a ArticleModel) SetConn(conn *sql.DB) ArticleModel {
	a.Conn = conn
	return a
}