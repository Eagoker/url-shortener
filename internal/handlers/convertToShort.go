package handlers

import (
	"io"
	"net/http"

	"github.com/Eagoker/url-shortener/pkg"

	"github.com/labstack/echo/v4"

)


func ConvertToShort(c echo.Context) error {
	if c.Request().Method != http.MethodPost {
		return echo.NewHTTPError(http.StatusMethodNotAllowed, "Method not allowed!")
	}

	originalURLBytes, err := io.ReadAll(c.Request().Body)
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, "Reading body error!")
	}
	
	originalURL := string(originalURLBytes)

	shortURL, err := pkg.GenerateShortURL(originalURL)
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, "Error with genrate short URL")
	}

	byteShortURL := []byte(shortURL)

	c.Response().Header().Set("content-type", "text/plain")
	return c.String(http.StatusCreated, string(byteShortURL))

}