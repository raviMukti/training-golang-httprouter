package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// Method Get
	router.GET("/", func(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(rw, "Hello HttpRouter")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	server.ListenAndServe()
}
