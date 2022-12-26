package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/hello", nil)
	helloHandler(recorder, request)

	body := recorder.Result().Body

	data, _ := io.ReadAll(body)

	assert.Equal(t, string(data), "Hello!")
}
