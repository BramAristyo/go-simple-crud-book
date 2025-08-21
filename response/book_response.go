package response

import "simple-restful-api/model/entity"

type BookResponse struct {
	Id     int
	Title  string
	Author string
	Year   int
}

func ToBookResponse(book entity.Book) BookResponse{
	return BookResponse{
		Id: book.Id,
		Title: book.Title,
		Author: book.Author,
		Year: book.Year,
	}
}

func ToBookResponses(books []entity.Book) []BookResponse{
	var bookResponses []BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, ToBookResponse(book))
	}

	return bookResponses
}