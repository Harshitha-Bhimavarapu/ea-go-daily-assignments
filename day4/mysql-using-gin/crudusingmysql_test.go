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

func makeGetApiCall(reqBody string) *httptest.ResponseRecorder {
	router := gin.New()

	router.GET("/books", GetBooks)
	request, _ := http.NewRequest("GET", "/books/", strings.NewReader(reqBody))
	responseWriter := httptest.NewRecorder()
	router.ServeHTTP(responseWriter, request)
	return responseWriter
}

func TestGetBooks(t *testing.T) {
	writer := makeGetApiCall("")

	responseBody, _ := io.ReadAll(writer.Body)
	response := Book{}
	json.Unmarshal(responseBody, &response)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "Dennis Ritchie", response.Author)
}
