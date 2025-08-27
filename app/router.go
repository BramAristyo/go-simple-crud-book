package app

import (
	"simple-restful-api/controller"
	"simple-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(bookController controller.BookController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/book",bookController.FindAll)
	router.GET("/api/book/:bookId",bookController.FindById)
	router.POST("/api/book", bookController.Create)
	router.PATCH("/api/book/:bookId", bookController.Update)
	router.DELETE("/api/book/:bookId", bookController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}