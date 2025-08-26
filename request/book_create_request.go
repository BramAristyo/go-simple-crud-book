package request

type BookCreateRequest struct {
	Title  string `json:"title" validate:"required,max=100,min=1"`
	Author string `json:"author" validate:"required"`
	Year   int    `json:"year" validate:"required,min=1"`
}