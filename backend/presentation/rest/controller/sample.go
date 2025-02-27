package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Sample interface {
	HelloWorld() echo.HandlerFunc
}

type sample struct{}

func NewSample() Sample {
	return &sample{}
}

func (s *sample) HelloWorld() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello World!"})
	}
}
