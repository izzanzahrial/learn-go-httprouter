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

// Not Found Handler
// by default we already have not found handler if we go to the undifened page
// but if want to change the not found page, you can use not found handler
func TestNotFoundHandler(t *testing.T) {
	router := httprouter.New()

	// Not found handler
	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Not Found")
	})

	// doesn't include this to make sure, every page is not found
	// router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// 	fmt.Fprint(writer, "Hello httprouter")
	// })

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil) // <- trying to access undifened page
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Not Found", string(body))
}
