package controller

import (
	"mvc/config"
	"mvc/model"
	"net/http"

	"github.com/labstack/echo"
)

//controller
func GetUserController(c echo.Context) error {
	var users []model.User
	err := config.DB.Find(&users).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}
func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)
	err := config.DB.Save(&user).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes create user",
		"user":    user,
	})

}
