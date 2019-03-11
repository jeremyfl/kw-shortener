package main

import (
	"kw-shortener/config"
	"kw-shortener/handlers"
	"kw-shortener/models"

	"github.com/labstack/echo"
)

func main() {
	config.Database().AutoMigrate(&models.Reference)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello World!")
	})

	e.GET("/short", handlers.GetAllUrl)

	e.Logger.Fatal(e.Start(":1323"))
}
