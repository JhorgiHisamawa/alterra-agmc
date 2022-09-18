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
	helpers.PanicIfError(err)
	if len(*res) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": helpers.EmptyData,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": helpers.Success,
		"data":    res,
	})
}

func GetOneUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	helpers.PanicIfError(err)

	data, err := database.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": helpers.UserNotFound,
			"data":    nil,
		})
	}
	res := helpers.UserResponse(*data)
	return c.JSON(http.StatusOK, res)
}

func CreateUser(c echo.Context) error {
	data := new(models.User)

	req, err := helpers.UserRequest(data, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": helpers.ErrorInput,
			"data":    nil,
		})
	}

	data, err = database.CreateUser(*req)
	helpers.PanicIfError(err)

	res := helpers.UserResponse(*data)
	return c.JSON(http.StatusCreated, res)
}

func UpdateUser(c echo.Context) error {
	d := new(models.User)
	id, err := strconv.Atoi(c.Param("id"))
	helpers.PanicIfError(err)
	_, err = database.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": helpers.UserNotFound,
		})
	}
	//check id is his or her own
	// userId := middlewares.ExtractTokenUserID(c)
	// if float64(id) != userId {
	// 	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
	// 		"message": helpers.ErrorUnauthorized,
	// 	})
	// }

	req, err := helpers.UserRequest(d, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": helpers.ErrorInput,
			"data":    nil,
		})
	}

	req.UpdatedAt = time.Now()
	data, err := database.UpdateUser(id, *req)
	helpers.PanicIfError(err)

	res := helpers.UserResponse(*data)
	return c.JSON(http.StatusOK, res)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	helpers.PanicIfError(err)

	//check id is his or her own
	// userId := middlewares.ExtractTokenUserID(c)
	// if float64(id) != userId {
	// 	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
	// 		"message": helpers.ErrorUnauthorized,
	// 	})
	// }

	getOne, err := database.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": helpers.UserNotFound,
		})
	}

	getOne.DeletedAt = time.Now()

	delAt, err := database.UpdateUser(id, *getOne)
	helpers.PanicIfError(err)

	_, err = database.DeleteUser(id)
	helpers.PanicIfError(err)

	res := helpers.UserResponse(*delAt)
	return c.JSON(http.StatusOK, res)
}
