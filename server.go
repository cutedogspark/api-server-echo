package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Working....")
	})

	e.GET("/ping",func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":    "pong",
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}