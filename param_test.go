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

func TestParam(t *testing.T) {
	router := httprouter.New()
	// the ":"id part is the param part
	router.GET("/products/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		productId := "Product " + params.ByName("id") // get the parameter value
		fmt.Fprint(writer, productId)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/1", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Product 1", string(body))
}

func TestNamedParam(t *testing.T) {
	router := httprouter.New()
	// You can have multiple named parameter
	router.GET("/products/:id/date/:date", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		productId := params.ByName("id")
		date := params.ByName("date")
		response := "Product " + productId + " at " + date
		fmt.Fprint(writer, response)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/1/date/04-04-2021", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Product 1 at 04-04-2021", string(body))
}

func TestAllParam(t *testing.T) {
	router := httprouter.New()
	// the "*" is to catch all the parameter in it
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		fmt.Fprint(writer, image)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/images/test_picture.jpeg", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "/test_picture.jpeg", string(body))
}
