package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", getHello)

	http.ListenAndServe(":3333", nil)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, HTTP!\n")
}