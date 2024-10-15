package pkg

import (
    "crypto/rand"
    "encoding/base64"
)

// GenerateShortURL принимает полный URL и возвращает короткий URL с рандомным концом.
func GenerateShortURL(fullURL string) (string, error) {
    // Генерируем случайную строку байт длиной 6 байт.
    randomBytes := make([]byte, 6)
    if _, err := rand.Read(randomBytes); err != nil {
        return "", err
    }

    // Кодируем байты в строку base64 и удаляем символы, не подходящие для URL.
    shortURL := base64.RawURLEncoding.EncodeToString(randomBytes)

    return shortURL, nil
}
