package pkg

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateShortURL принимает полный URL и возвращает короткий ID.
func GenerateShortURL() (string, error) {
    // Генерируем случайную строку байт длиной 6 байт.
    randomBytes := make([]byte, 6)
    _, err := rand.Read(randomBytes)
    if err != nil {
        return "", err
    }

    // Кодируем байты в строку base64 и удаляем символы, не подходящие для URL.
    shortID := base64.RawURLEncoding.EncodeToString(randomBytes)

    // Возвращаем только короткий ID
    return shortID, nil
}
