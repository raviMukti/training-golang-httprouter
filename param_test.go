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

	router.GET("/products/:id", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		text := "Product " + id
		fmt.Fprint(rw, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1", string(bytes))
}
