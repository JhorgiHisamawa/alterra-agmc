package controllers

import (
	"api/internals/models"
	"api/internals/repositories"
	"api/pkg/util"
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	req := &models.User{}

	req, err := util.UserRequest(req, c)
	util.PanicIfError(err)

	data, err := repositories.LoginRepository(req)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
		})
	}

	res := util.UserResponse(*data)

	return c.JSON(http.StatusOK, res)
}

func Register(c echo.Context) error {
	data := new(models.User)

	req, err := util.UserRequest(data, c)
	util.PanicIfError(err)

	data, err = repositories.RegisterRepository(req)
	util.PanicIfError(err)

	res := util.UserResponse(*data)

	return c.JSON(http.StatusCreated, res)
}
