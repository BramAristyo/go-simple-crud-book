package controller

import (
	"net/http"
	"simple-restful-api/helper"
	"simple-restful-api/request"
	"simple-restful-api/response"
	"simple-restful-api/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func (controller *BookControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookCreateRequest := request.BookCreateRequest{}
	helper.ReadRequestBody(r, &bookCreateRequest)
	bookResponse := controller.BookService.Create(r.Context(), bookCreateRequest)

	webResponse := response.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bookResponse,
	}

	helper.WriteToResponse(w, webResponse)
}

func (controller *BookControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookUpdateRequest := request.BookUpdateRequest{}
	helper.ReadRequestBody(r, &bookUpdateRequest)

	bookId := p.ByName("bookId")
	Id, err := strconv.Atoi(bookId)
	helper.PanicIfErr(err)

	bookUpdateRequest.Id = Id

	bookResponse := controller.BookService.Update(r.Context(), bookUpdateRequest)

	webResponse := response.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bookResponse,
	}

	helper.WriteToResponse(w, webResponse)
}

func (controller *BookControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookId := p.ByName("bookId")
	Id, err := strconv.Atoi(bookId)
	helper.PanicIfErr(err)

	controller.BookService.Delete(r.Context(), Id)

	webResponse := response.WebResponse{
		Code: 200,
		Status: "OK",
	}

	helper.WriteToResponse(w, webResponse)
}

func (controller *BookControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookId := p.ByName("bookId")
	Id, err := strconv.Atoi(bookId)
	helper.PanicIfErr(err)

	bookResponse := controller.BookService.FindById(r.Context(), Id)

	webResponse := response.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bookResponse,
	}

	helper.WriteToResponse(w, webResponse)
}

func (controller *BookControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookResponses := controller.BookService.FindAll(r.Context())

	webResponse := response.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bookResponses,
	}

	helper.WriteToResponse(w, webResponse)
}