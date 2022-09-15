package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func LogMiddleware(c *echo.Echo) {
	c.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, status=${status}, uri=${uri} \n",
	}))
}
