package request

type BookUpdateRequest struct {
	Id     int
	Title  string
	Author string
	Year   int
}