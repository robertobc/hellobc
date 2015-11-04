package main

import (
	"fmt"
	"net/http"

	"github.com/robertobc/hellobc/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Handle)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
