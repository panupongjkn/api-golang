package routes

import (
	"api-golang/action/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserRoutes(e *echo.Echo) {
	g := e.Group("/user")
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("golang"),
		// TokenLookup: "token",
	}))
	g.GET("", controllers.GetUser)
	e.POST("/signin", controllers.SignIn)
	e.POST("/signup", controllers.SignUp)
}
