package handlers

import (
	"fmt"
	"net/http"
)

type HomeHandler struct{}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a basic Home Page....")
}
