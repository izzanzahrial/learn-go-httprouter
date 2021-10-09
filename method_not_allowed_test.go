package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// Method not allowed
// by default we already have method not allowed if use the wrong method when trying to access the page
// but if want to change method not allowed, we can use
func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()

	// Method not allowed
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Not Allowed unless GET METHOD")
	})

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello httprouter")
	})

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil) // <- trying to access using POST method
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Not Allowed unless GET METHOD", string(body))
}
