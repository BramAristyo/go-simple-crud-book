package main

import (
	"fmt"
	"net/http"
	"simple-restful-api/config"
	"simple-restful-api/controller"
	"simple-restful-api/exception"
	"simple-restful-api/helper"
	"simple-restful-api/middleware"
	"simple-restful-api/repository"
	"simple-restful-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := config.NewDB()
	validate := validator.New()
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, db, validate)
	bookController := controller.NewBookController(bookService)

	router := httprouter.New()

	router.GET("/api/book",bookController.FindAll)
	router.GET("/api/book/:bookId",bookController.FindById)
	router.POST("/api/book", bookController.Create)
	router.PATCH("/api/book/:bookId", bookController.Update)
	router.DELETE("/api/book/:bookId", bookController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: "localhost:8000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Server successfully connected to port 8000")
	err := server.ListenAndServe()
	helper.PanicIfErr(err)

}