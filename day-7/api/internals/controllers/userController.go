package controllers

import (
	"api/internals/middlewares"
	"api/internals/models"
	database "api/internals/repositories"
	"api/pkg/util"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func GetAllUser(c echo.Context) error {
	res, err := database.GetUser()
	util.PanicIfError(err)
	if len(*res) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": util.EmptyData,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": util.Success,
		"data":    res,
	})
}

func GetOneUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	util.PanicIfError(err)

	data, err := database.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": util.UserNotFound,
			"data":    nil,
		})
	}
	res := util.UserResponse(*data)
	return c.JSON(http.StatusOK, res)
}

func CreateUser(c echo.Context) error {
	data := new(models.User)

	req, err := util.UserRequest(data, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": util.ErrorInput,
			"data":    nil,
		})
	}

	data, err = database.CreateUser(*req)
	util.PanicIfError(err)

	res := util.UserResponse(*data)
	return c.JSON(http.StatusCreated, res)
}

func UpdateUser(c echo.Context) error {
	d := new(models.User)
	id, err := strconv.Atoi(c.Param("id"))
	util.PanicIfError(err)
	_, err = database.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": util.UserNotFound,
		})
	}
	//check id is his or her own
	userId := middlewares.ExtractTokenUserID(c)
	if float64(id) != userId {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": util.ErrorUnauthorized,
		})
	}

	req, err := util.UserRequest(d, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": util.ErrorInput,
			"data":    nil,
		})
	}

	req.UpdatedAt = time.Now()
	data, err := database.UpdateUser(id, *req)
	util.PanicIfError(err)

	res := util.UserResponse(*data)
	return c.JSON(http.StatusOK, res)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	util.PanicIfError(err)

	//check id is his or her own
	userId := middlewares.ExtractTokenUserID(c)
	if float64(id) != userId {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": util.ErrorUnauthorized,
		})
	}

	getOne, err := database.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": util.UserNotFound,
		})
	}

	getOne.DeletedAt = time.Now()

	delAt, err := database.UpdateUser(id, *getOne)
	util.PanicIfError(err)

	_, err = database.DeleteUser(id)
	util.PanicIfError(err)

	res := util.UserResponse(*delAt)
	return c.JSON(http.StatusOK, res)
}
