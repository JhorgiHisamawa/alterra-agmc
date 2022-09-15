package controllers

import (
	"api/helpers"
	"api/lib/database"
	"api/models"
	"net/http"

	"github.com/labstack/echo"
)

func LoginController(c echo.Context) error {
	req := &models.User{}

	req, err := helpers.UserRequest(req, c)
	helpers.PanicIfError(err)

	data, err := database.LoginRepository(req)
	helpers.PanicIfError(err)

	res := helpers.UserResponse(*data)

	return c.JSON(http.StatusOK, res)
}

func RegisterController(c echo.Context) error {
	data := new(models.User)

	req, err := helpers.UserRequest(data, c)
	helpers.PanicIfError(err)

	data, err = database.RegisterRepository(req)
	helpers.PanicIfError(err)

	res := helpers.UserResponse(*data)

	return c.JSON(http.StatusCreated, res)
}
