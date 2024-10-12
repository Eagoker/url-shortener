package handlers

import (
	"net/http"

	"github.com/Eagoker/url-shortener/pkg"

	"github.com/labstack/echo/v4"
)

type Request struct {
	OriginalUrl string `json:"url"`
}

type Responce struct {
	ShortUrl string `json:"result"`
}

func ApiShorten(c echo.Context) error {
	if c.Request().Method != http.MethodPost{
		return echo.NewHTTPError(http.StatusMethodNotAllowed, "Method not allowed!")
	}

	requestBody := new(Request)

	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	shortURL, err := pkg.GenerateShortURL(requestBody.OriginalUrl)
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, "Error with genrate short URL")
	}

	responceBody := Responce{ShortUrl: shortURL}
	
	return c.JSON(http.StatusCreated, responceBody)
}