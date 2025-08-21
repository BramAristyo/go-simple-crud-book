package request

type BookUpdateRequest struct {
	Id     int    `validate:"required"`
	Title  string `validate:"required,max=100,min=1"`
	Author string `validate:"required"`
	Year   int    `validate:"required,min=1,max=4"`
}