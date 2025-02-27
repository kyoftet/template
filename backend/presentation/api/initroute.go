package api

import (
	"backend/presentation/api/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(
	userRouter router.User,
	authRouter router.Auth,
	sampleRouter router.Sample,
) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", authRouter.Login())

	e.POST("/signup", userRouter.SignUp())

	authGroup := e.Group("/")
	authGroup.Use(authRouter.Middleware)

	authGroup.POST("/logout", authRouter.Logout())

	authGroup.DELETE("/delete", userRouter.Delete())
	authGroup.PUT("/profile", userRouter.Update())

	authGroup.GET("/hello-world", sampleRouter.HelloWorld())

	e.Logger.Fatal(e.Start(":8000"))
}
