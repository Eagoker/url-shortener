package handlers

import (
	"net/http"
	"strings"
)

func GetOriginalUrl(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
	}
	
	path := strings.TrimPrefix(r.URL.Path, "/") // Убираем первый слеш
    id := strings.Split(path, "/")[0] 

	//тут будет логика получения ориг юрла
	_ = id

	originalUrl := "https://practicum.yandex.ru/"
	
	w.Header().Set("Location", originalUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
}