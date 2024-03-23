package main

import (
	"demo_api/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &handlers.HomeHandler{})
	mux.Handle("/books/", &handlers.BookHandler{})

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		return
	}
}
