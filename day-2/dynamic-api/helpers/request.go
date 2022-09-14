package helpers

import (
	"api/models"

	"github.com/labstack/echo"
)

func Request(data *models.User, c echo.Context) (*models.User, error) {
	c.Request().Header.Get("Content-type")
	err := c.Bind(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
