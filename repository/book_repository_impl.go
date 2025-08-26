package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple-restful-api/helper"
	"simple-restful-api/model/entity"
)

type BookRepositoryImpl struct {
	
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (repository *BookRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, book entity.Book) entity.Book {
	q := "INSERT INTO books(title, author, year) VALUES (?, ?, ?)"
	res, err := tx.ExecContext(ctx, q, book.Title, book.Author, book.Year)
	helper.PanicIfErr(err)

	id, err := res.LastInsertId()
	helper.PanicIfErr(err)

	book.Id = int(id)
	return book
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book entity.Book) entity.Book {
    if tx == nil {
        panic(errors.New("transaction is nil"))
    }
    q := "UPDATE books SET title = ?, author = ?, year = ? WHERE id = ?"
    _, err := tx.ExecContext(ctx, q, book.Title, book.Author, book.Year, book.Id)
    helper.PanicIfErr(err)
    return book
}

func (repository *BookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, bookId int){
	q := "DELETE FROM books where id = ?"
	_, err := tx.ExecContext(ctx, q, bookId)
	helper.PanicIfErr(err)
}

func (repository *BookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookId int) (entity.Book, error) {
	q := "SELECT id, title, author, year FROM books where id = ?"
	rows, err := tx.QueryContext(ctx, q, bookId)
	helper.PanicIfErr(err)
	defer rows.Close()

	book := entity.Book{}
	if rows.Next() {
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Year)
		helper.PanicIfErr(err)
		return book, nil
	} else {
		return book, errors.New("book is not found")
	}

}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Book {
	q := "SELECT id, title, author, year FROM books"
	rows, err := tx.QueryContext(ctx, q)
	helper.PanicIfErr(err)
	defer rows.Close()

	var books []entity.Book
	for rows.Next() {
		book := entity.Book{}
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Year)
		helper.PanicIfErr(err)
		
		books = append(books, book)
	}

	return books
}



