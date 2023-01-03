package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router = gin.New()

func makeGetCall(reqBody string) *httptest.ResponseRecorder {
	router.GET("/books", getBooks)
	request, _ := http.NewRequest("GET", "/books", strings.NewReader(reqBody))
	responseWriter := httptest.NewRecorder()
	router.ServeHTTP(responseWriter, request)
	return responseWriter
}

func makePostCall(reqBody string) *httptest.ResponseRecorder {
	router.POST("/books", createBook)
	request, _ := http.NewRequest("POST", "/books", strings.NewReader(reqBody))
	responseWriter := httptest.NewRecorder()
	router.ServeHTTP(responseWriter, request)
	return responseWriter
}

func makeUpdateCall(reqBody string) *httptest.ResponseRecorder {
	router.PUT("/books/:Id", updateBook)
	request, _ := http.NewRequest("PUT", "/books/Id", strings.NewReader(reqBody))
	responseWriter := httptest.NewRecorder()
	router.ServeHTTP(responseWriter, request)
	return responseWriter
}

func TestGetBooks(t *testing.T) {
	reqBody := `{
			"Id": 2,
			"Author": "William Stallings",
			"Price": 56.99
	}`

	writer := makeGetCall(reqBody)

	assert.Equal(t, http.StatusOK, writer.Code)
	respBody, _ := io.ReadAll(writer.Body)
	response := Book{}
	json.Unmarshal(respBody, &response)
	assert.Equal(t, 1, response.Id)
	assert.Equal(t, "William Stallings", response.Author)

}

func TestCreateBook(t *testing.T) {
	reqBody := `{
		"Id":5,
		"Author":"Tenananbaum",
		"Price":100.00
	}`

	writer := makePostCall(reqBody)

	respBody, _ := io.ReadAll(writer.Body)
	response := Book{}
	json.Unmarshal(respBody, &response)

	assert.Equal(t, http.StatusCreated, writer.Code)
	assert.Equal(t, "Tenananbaum", response.Author)
}

func TestUpdateBook(t *testing.T) {
	reqBody := `{
		"Id":1,
		"Author":"James Gosling",
		"Price":456.89
	}`

	writer := makeUpdateCall(reqBody)
	response := Book{}
	resBody, _ := io.ReadAll(writer.Body)
	json.Unmarshal(resBody, &response)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "James Gosling", response.Author)
	assert.Equal(t, 456.89, response.Price)
}
