package main

import (
	routes "myapp/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})
	uGroup := e.Group("/users")
	routes.NewUserRoute(uGroup)

	e.Logger.Fatal(e.Start(":8080"))
}
