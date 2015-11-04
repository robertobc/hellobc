package main

import "net/http"
import "./handlers"

func main() {
	http.HandleFunc("/", handlers.Handle)
	http.ListenAndServe(":8080", nil)
}
