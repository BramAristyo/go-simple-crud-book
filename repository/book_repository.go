package repository

import (
	"context"
	"database/sql"
	"simple-restful-api/model/entity"
)

type BookRepository interface {
	Save(ctx context.Context, tx *sql.Tx, book entity.Book) entity.Book
	Update(ctx context.Context, tx *sql.Tx, book entity.Book) entity.Book
	Delete(ctx context.Context, tx *sql.Tx, bookId int) bool
	FindById(ctx context.Context, tx *sql.Tx, bookId int) (entity.Book, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Book
}