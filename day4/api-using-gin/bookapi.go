package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     int     `json:"Id"`
	Author string  `json:"Author"`
	Price  float64 `json:"Price"`
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.PUT("/books/Id", updateBook)

	router.Run("localhost:8080")
}

func getBooks(ctx *gin.Context) {
	request := Book{}
	ctx.BindJSON(&request)
	response := Book{
		Id:     1,
		Author: request.Author,
		Price:  request.Price,
	}
	ctx.JSON(http.StatusOK, response)
}

func createBook(ctx *gin.Context) {
	var input Book
	ctx.ShouldBindJSON(&input)
	response := Book{
		Id:     input.Id,
		Author: input.Author,
		Price:  input.Price,
	}
	ctx.JSON(http.StatusCreated, response)
}

func updateBook(ctx *gin.Context) {
	request := Book{}
	ctx.BindJSON(&request)
	response := Book{
		Id:     request.Id,
		Author: request.Author,
		Price:  request.Price,
	}
	ctx.JSON(http.StatusOK, response)
}
