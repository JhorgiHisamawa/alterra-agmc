package helpers

import (
	"api/models"

	"github.com/labstack/echo"
)

func Request(data models.Books, c echo.Context) (*models.Books, error) {
	c.Request().Header.Get("Content-type")
	err := c.Bind(&data)
	if err != nil {
		return &data, err
	}
	return &data, nil
}
