package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Book struct {
	ID     uint    `json:”id”`
	Author string  `json:author`
	Price  float64 `json:price`
}

func main() {
	SetUpDB()
	defer DB.Close()

	book1 := Book{2, "William stallings", float64(1234.67)}
	book2 := Book{3, "Dennis Ritchie", float64(1834.67)}

	DB.Create(&book1)
	DB.Create(book2)

	router := gin.Default()
	router.GET("/books", GetBooks)
	router.GET("/books/:id", GetBook)
	router.POST("/books", CreateBook)
	router.PUT("/books/:id", UpdateBook)
	router.DELETE("/books/:id", DeleteBook)
	router.Run("localhost:8080")
}

func GetBooks(ctx *gin.Context) {
	var books []Book
	if err := DB.Find(&books).Error; err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(http.StatusOK, books)
	}
}

func GetBook(ctx *gin.Context) {
	var book Book
	id := ctx.Params.ByName("id")
	if err := DB.Where("id=?", id).First(&book).Error; err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(http.StatusOK, book)
	}
}

func CreateBook(ctx *gin.Context) {
	var book Book
	ctx.BindJSON(&book)
	DB.Create(&book)
	ctx.JSON(http.StatusOK, book)
}

func UpdateBook(ctx *gin.Context) {
	var book Book
	id := ctx.Params.ByName("id")
	if err := DB.Where("id=?", id).Error; err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	}
	ctx.BindJSON(&book)
	DB.Save(&book)
	ctx.JSON(http.StatusOK, book)
}

func DeleteBook(ctx *gin.Context) {
	var book Book
	id := ctx.Params.ByName("id")
	d := DB.Where("id = ?", id).Delete(&book)
	fmt.Println(d)
	ctx.JSON(200, gin.H{"id #" + id: "deleted"})
}
