package controllers

import (
	"api/helpers"
	"api/lib/database"
	"api/models"
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	req := &models.User{}

	req, err := helpers.UserRequest(req, c)
	helpers.PanicIfError(err)

	data, err := database.LoginRepository(req)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
		})
	}

	res := helpers.UserResponse(*data)

	return c.JSON(http.StatusOK, res)
}

func Register(c echo.Context) error {
	data := new(models.User)

	req, err := helpers.UserRequest(data, c)
	helpers.PanicIfError(err)

	data, err = database.RegisterRepository(req)
	helpers.PanicIfError(err)

	res := helpers.UserResponse(*data)

	return c.JSON(http.StatusCreated, res)
}
