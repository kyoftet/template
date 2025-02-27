package router

import (
	"backend/application/dto"
	"backend/application/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User interface {
	SignUp() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type user struct {
	applicationService service.User
}

func NewUser(applicationService service.User) User {
	return &user{applicationService: applicationService}
}

func (u *user) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(dto.User)
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
		}
		err := u.applicationService.SignUp(user)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Could not create user"},
			)
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "User created successfully"})
	}
}

func (u *user) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userId string
		if err := c.Bind(userId); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
		}
		err := u.applicationService.Delete(userId)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Could not create user"},
			)
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "User created successfully"})
	}
}

func (u *user) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(dto.User)
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
		}
		err := u.applicationService.Update(user)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"message": "Could not create user"},
			)
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "User created successfully"})
	}
}
