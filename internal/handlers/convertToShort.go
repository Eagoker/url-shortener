package handlers

import (
	"io"
	"net/http"

	"github.com/Eagoker/url-shortener/pkg"
)

func ConvertToShort(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	originalUrlBytes, err := io.ReadAll(r.Body)
	if err != nil{
		http.Error(w, "Error with reading body!", http.StatusBadRequest)	
	}

	originalUrl := string(originalUrlBytes)

	shortUrl, err := pkg.GenerateShortURL(originalUrl)
	if err != nil{
		http.Error(w, "Error with generating shortURL!", http.StatusInternalServerError)
	}

	byteShortUrl := []byte(shortUrl)

	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write(byteShortUrl)
}