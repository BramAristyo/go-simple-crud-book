package service

import (
	"context"
	"database/sql"
	"simple-restful-api/exception"
	"simple-restful-api/helper"
	"simple-restful-api/model/entity"
	"simple-restful-api/repository"
	"simple-restful-api/request"
	"simple-restful-api/response"

	"github.com/go-playground/validator/v10"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewBookService(bookRepository repository.BookRepository, DB *sql.DB, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
		DB: DB,
		Validate: validate,
	}
}

func (service *BookServiceImpl) Create(ctx context.Context, request request.BookCreateRequest) response.BookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollBack(tx)

	book := entity.Book{
		Title: request.Title,
		Author: request.Author,
		Year: request.Year,
	}

	book = service.BookRepository.Save(ctx, tx, book)
	return response.ToBookResponse(book)
}

func (service *BookServiceImpl) Update(ctx context.Context, request request.BookUpdateRequest) response.BookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)
	
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollBack(tx)

	_, err = service.BookRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	book := entity.Book{
		Id: request.Id,
		Title: request.Title,
		Author: request.Author,
		Year: request.Year,
	}

	book = service.BookRepository.Update(ctx, tx, book)
	return response.ToBookResponse(book)
}

func (service *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollBack(tx)

	_, err = service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.BookRepository.Delete(ctx, tx, bookId)
}

func (service *BookServiceImpl) FindById(ctx context.Context, bookId int) response.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollBack(tx)

	book, err:= service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return response.ToBookResponse(book)
}

func (service *BookServiceImpl) FindAll(ctx context.Context) []response.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollBack(tx)

	books := service.BookRepository.FindAll(ctx, tx)
	return response.ToBookResponses(books)
}

