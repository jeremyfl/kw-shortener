package main

import (
	"kw-shortener/config"
	"kw-shortener/handlers"
	"kw-shortener/helpers"
	"kw-shortener/models"

	"github.com/labstack/echo"
)

func main() {
	config.Database().AutoMigrate(&models.Reference{})
	config.Database().Close()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, helpers.Response{Message: "Hello world!"})
	})

	e.GET("/proto", DecodeProto)

	e.GET("/reference", handlers.GetAllReference)
	e.GET("/reference/:id", handlers.ShowReference)
	e.POST("/reference", handlers.StoreReference)

	e.Logger.Fatal(e.Start(":1323"))
}
