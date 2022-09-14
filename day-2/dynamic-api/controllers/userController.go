package controllers

import (
	"api/helpers"
	"api/lib/database"
	"api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func GetAllUser(c echo.Context) error {
	res, err := database.GetUser()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "data is empty",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetOneUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	helpers.PanicIfError(err)

	res, err := database.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func CreateUser(c echo.Context) error {
	data := new(models.User)

	req, err := helpers.Request(data, c)
	helpers.PanicIfError(err)

	data, err = database.CreateUser(*req)
	helpers.PanicIfError(err)

	res := helpers.Response(*data)

	return c.JSON(http.StatusCreated, res)
}

func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	d := new(models.User)

	req, err := helpers.Request(d, c)
	helpers.PanicIfError(err)

	req.UpdatedAt = time.Now()

	data, err := database.UpdateUser(id, *req)
	helpers.PanicIfError(err)

	res := helpers.Response(*data)

	return c.JSON(http.StatusOK, res)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	helpers.PanicIfError(err)

	getOne, err := database.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
		})
	}

	getOne.DeletedAt = time.Now()

	delAt, err := database.UpdateUser(id, *getOne)
	helpers.PanicIfError(err)

	_, err = database.DeleteUser(id)
	helpers.PanicIfError(err)

	res := helpers.Response(*delAt)

	return c.JSON(http.StatusOK, res)
}
