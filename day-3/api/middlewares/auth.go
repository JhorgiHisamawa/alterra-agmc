package middlewares

import (
	"api/config"

	"github.com/labstack/echo"
)

const (
	ADMIN = "admin"
)

func BasicAuth(email, password, role string, c echo.Context) (bool, error) {
	var err error
	if email == config.Config("EMAIL") && password == config.Config("PASSWORD") {
		return true, err
	}
	return false, nil
}
