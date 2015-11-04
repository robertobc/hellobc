package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

//This function handles requests
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Blabcake :)")
}
