package main

import (
	"fmt"
	"net/http"
	"simple-restful-api/app"
	"simple-restful-api/controller"
	"simple-restful-api/helper"
	"simple-restful-api/middleware"
	"simple-restful-api/repository"
	"simple-restful-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, db, validate)
	bookController := controller.NewBookController(bookService)
	router := app.NewRouter(bookController)

	server := http.Server{
		Addr: "localhost:8000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Server successfully connected to port 8000")
	err := server.ListenAndServe()
	helper.PanicIfErr(err)

}