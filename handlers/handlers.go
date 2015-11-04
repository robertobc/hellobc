package handlers

import (
	"fmt"
	"net/http"
)

//This function handles requests
func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Blabcake :)")
}
