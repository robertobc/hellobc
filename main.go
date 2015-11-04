package main

import "net/http"
import "github.com/robertobc/hellobc/handlers"

func main() {
	http.HandleFunc("/", handlers.Handle)
	http.ListenAndServe(":8080", nil)
}
