package main

import (
	"andromeda/belajar-golang-restful-api/app"
	"andromeda/belajar-golang-restful-api/controller"
	"andromeda/belajar-golang-restful-api/helper"
	"andromeda/belajar-golang-restful-api/middleware"
	"andromeda/belajar-golang-restful-api/repository"
	"andromeda/belajar-golang-restful-api/service"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/go-playground/validator/v10"
)

func main() {

	err := godotenv.Load()
	helper.PanicIfError(err)

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:" + os.Getenv("PORT"),
		Handler: middleware.NewAuthMiddleware(router),
	}

	log.Println("API is running!")
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
