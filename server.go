package main

import (
	"kw-shortener/config"
	"kw-shortener/handlers"
	"kw-shortener/models"

	"github.com/labstack/echo"
)

func main() {
	config.Database().AutoMigrate(&models.Reference{})
	config.Database().Close()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello World!")
	})

	e.GET("/reference", handlers.GetAllReference)
	e.POST("/reference", handlers.InsertReference)

	e.Logger.Fatal(e.Start(":1323"))
}
