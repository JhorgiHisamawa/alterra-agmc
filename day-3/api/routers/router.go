package routers

import (
	"api/config"
	"api/controllers"
	m "api/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)
	// Route / to handler function

	r := e.Group("/v1")
	r.POST("/login", controllers.LoginController)
	r.POST("/register", controllers.RegisterController)

	r.POST("/users", controllers.CreateUser)
	// jwt token
	r.GET("/users", controllers.GetAllUser, middleware.JWT([]byte(config.Config("SECRET_JWT"))))
	r.GET("/users/:id", controllers.GetOneUser, middleware.JWT([]byte(config.Config("SECRET_JWT"))))
	r.PUT("/users/:id", controllers.UpdateUser, middleware.JWT([]byte(config.Config("SECRET_JWT"))))
	r.DELETE("/users/:id", controllers.DeleteUser, middleware.JWT([]byte(config.Config("SECRET_JWT"))))

	r.GET("/books", controllers.GetAllBook)
	r.GET("/books/:id", controllers.GetOneBook)
	// basic auth
	r.POST("/books", controllers.CreateBook, middleware.BasicAuth(m.BasicAuthDB))
	r.PUT("/books/:id", controllers.UpdateBook, middleware.BasicAuth(m.BasicAuthDB))
	r.DELETE("/books/:id", controllers.DeleteBook, middleware.BasicAuth(m.BasicAuthDB))

	return e
}
