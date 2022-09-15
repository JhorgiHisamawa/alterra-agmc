package helpers

import (
	"api/models"

	"github.com/labstack/echo"
)

func UserRequest(data *models.User, c echo.Context) (*models.User, error) {
	c.Request().Header.Get("Content-type")
	err := c.Bind(&data)
	if err != nil {
		return data, err
	}
	//validate data
	if err = c.Validate(data); err != nil {
		return data, err
	}
	return data, nil
}

func BookRequest(data *models.Books, c echo.Context) (*models.Books, error) {
	c.Request().Header.Get("Content-type")
	err := c.Bind(&data)
	if err != nil {
		return data, err
	}
	//validate data
	if err = c.Validate(data); err != nil {
		return data, err
	}
	return data, nil
}
