package pkg

import (
    "crypto/rand"
    "encoding/base64"
    "fmt"
)

// generateShortURL принимает полный URL и возвращает короткий URL с рандомным концом.
func GenerateShortURL(fullURL string) (string, error) {
    // Генерируем случайную строку байт длиной 6 байт.
    randomBytes := make([]byte, 6)
    _, err := rand.Read(randomBytes)
    if err != nil {
        return "", err
    }

    // Кодируем байты в строку base64 и удаляем символы, не подходящие для URL.
    shortID := base64.RawURLEncoding.EncodeToString(randomBytes)

    // Формируем короткий URL.
    shortURL := fmt.Sprintf("http://localhost:8080/%s", shortID)
    return shortURL, nil
}
