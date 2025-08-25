package config

import (
	"database/sql"
	"simple-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3307)/crud-book-go")
	helper.PanicIfErr(err)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetConnMaxIdleTime(time.Minute * 60)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}