package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (h *HTTPHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if len(body) == 0 {
			return c.String(http.StatusBadRequest, "You should set body")
		}

		urlIdentifier := uuid.New().String()
		shortURL := host + urlIdentifier
		h.urlStorage.Add(urlIdentifier, string(body))

		return c.String(http.StatusCreated, shortURL)
	}
}
