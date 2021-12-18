package route

import (
	"mvc/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	return e
}
