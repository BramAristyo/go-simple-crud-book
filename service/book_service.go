package service

import (
	"context"
	"simple-restful-api/request"
	"simple-restful-api/response"
)

type BookService interface {
	Create(ctx context.Context, request request.BookCreateRequest) response.BookResponse
	Update(ctx context.Context, request request.BookUpdateRequest) response.BookResponse
	Delete(ctx context.Context, bookId int) bool
	FindById(ctx context.Context, bookId int) response.BookResponse
	FindAll(ctx context.Context) []response.BookResponse
}