package request

type BookUpdateRequest struct {
	Id     int    `json:"id" validate:"required"`
	Title  string `json:"title" validate:"required,max=100,min=1"`
	Author string `json:"author" validate:"required"`
	Year   int    `json:"year" validate:"required"`
}