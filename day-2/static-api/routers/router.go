package routers

import (
	"api/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	// Route / to handler function
	g := e.Group("/v1")
	g.GET("/books", controllers.GetAllBook)
	g.GET("/books/:id", controllers.GetOneBook)
	g.POST("/books", controllers.CreateBook)
	g.PUT("/books/:id", controllers.UpdateBook)
	g.DELETE("/books/:id", controllers.DeleteBook)

	return e
}
