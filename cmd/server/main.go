package main

import (
	"net/http"

	"github.com/Eagoker/url-shortener/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.ConvertToShort)
	mux.HandleFunc("/{id}", handlers.GetOriginalUrl)

	if err := http.ListenAndServe(":8080", mux); err != nil{
		panic(err)
	}
	
}