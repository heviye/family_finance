package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		c.Logger().Info("print test log ... ")
		return c.String(http.StatusOK, "Hello Hevi")
	})

	err := e.Start(":8010")
	if err != nil {
		panic(err)
	}
}
