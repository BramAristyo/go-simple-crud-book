package config

import (
	"database/sql"
	"simple-restful-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3307)/crud-book-go")
	helper.PanicIfErr(err)

	return db
}