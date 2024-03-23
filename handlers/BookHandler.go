package handlers

import (
	"fmt"
	"net/http"
	"regexp"
)

type BookHandler struct{}

var (
	BookRe       = regexp.MustCompile(`^/books/?$`)
	BookReWithID = regexp.MustCompile(`^/books/([a-zA-Z0-9]+)$`)
)

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Creating a Book\n")
}
func (h *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Listing all the books\n")
}
func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprintf(w, "Getting a book with id: %s\n", id)
}
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprintf(w, "Updating a book with id: %s\n", id)
}
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprintf(w, "Deleting a bool with id: %s\n", id)
}

func (h *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && BookRe.MatchString(r.URL.Path):
		h.CreateBook(w, r)
		return
	case r.Method == http.MethodGet && BookRe.MatchString(r.URL.Path):
		h.ListBooks(w, r)
		return
	case r.Method == http.MethodGet && BookReWithID.MatchString(r.URL.Path):
		matches := BookReWithID.FindStringSubmatch(r.URL.Path)
		if len(matches) > 1 {
			id := matches[1]
			h.GetBook(w, r, id)
		} else {
			http.NotFound(w, r)
		}
	case r.Method == http.MethodPost && BookReWithID.MatchString(r.URL.Path):
		matches := BookReWithID.FindStringSubmatch(r.URL.Path)
		if len(matches) > 1 {
			id := matches[1]
			h.UpdateBook(w, r, id)
		} else {
			http.NotFound(w, r)
		}
	case r.Method == http.MethodDelete && BookReWithID.MatchString(r.URL.Path):
		matches := BookReWithID.FindStringSubmatch(r.URL.Path)
		if len(matches) > 1 {
			id := matches[1]
			h.DeleteBook(w, r, id)
		} else {
			http.NotFound(w, r)
		}
	default:
		http.NotFound(w, r)
	}
}
