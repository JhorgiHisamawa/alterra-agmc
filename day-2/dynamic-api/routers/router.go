package routers

import (
	"api/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	// Route / to handler function

	g := e.Group("/v1")
	g.GET("/users", controllers.GetAllUser)
	g.GET("/users/:id", controllers.GetOneUser)
	g.POST("/users", controllers.CreateUser)
	g.PUT("/users/:id", controllers.UpdateUser)
	g.DELETE("/users/:id", controllers.DeleteUser)

	return e
}
