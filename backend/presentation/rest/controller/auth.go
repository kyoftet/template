package controller

import (
	"backend/application/dto"
	"backend/application/service"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Auth interface {
	Login() echo.HandlerFunc
	Logout() echo.HandlerFunc
	Middleware(next echo.HandlerFunc) echo.HandlerFunc
}

type auth struct {
	applicatoinService service.Auth
}

func NewAuth(applicatoinService service.Auth) Auth {
	return &auth{applicatoinService: applicatoinService}
}

func (a *auth) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		credential := new(dto.Auth)
		if err := c.Bind(credential); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
		}

		sessionId, err := a.applicatoinService.Login(credential)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
		}

		c.SetCookie(
			&http.Cookie{
				Name:     "session_id",
				Value:    sessionId,
				Expires:  time.Now().Add(24 * time.Hour),
				HttpOnly: true,
			},
		)
		return c.JSON(http.StatusOK, map[string]string{"message": "Login successful"})
	}
}

func (a *auth) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session_id")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "No session ID found"})
		}

		sessionId := cookie.Value

		err = a.applicatoinService.Logout(sessionId)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Could not delete session"})
		}

		c.SetCookie(&http.Cookie{
			Name:     "session_id",
			Value:    "",
			Expires:  time.Now().Add(-1 * time.Hour),
			HttpOnly: true,
		})
		return c.JSON(http.StatusOK, map[string]string{"message": "Logout successful"})
	}
}

func (a *auth) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session_id")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing session ID"})
		}

		sessionId := cookie.Value
		userId, err := a.applicatoinService.GetSessionId(sessionId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error checking session ID"})
		}

		c.Set("userId", userId)
		return next(c)
	}
}
