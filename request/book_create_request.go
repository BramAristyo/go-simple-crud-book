package request

type BookCreateRequest struct {
	Title  string `validate:"required,max=100,min=1"`
	Author string `validate:"required"`
	Year   int    `validate:"required,min=1"`
}