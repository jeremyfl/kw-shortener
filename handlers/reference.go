package handlers

import (
	"kw-shortener/models"

	"github.com/labstack/echo"
)

// GetAllReference list of all url reference
func GetAllReference(c echo.Context) error {
	return c.JSON(200, "All Reference goes here")
}

func InsertReference(c echo.Context) (err error) {
	r := new(models.Reference)

	if err = c.Bind(r); err != nil {
		return c.JSON(400, "Bad Request")
	}

	insertReference := models.StoreReference(r)

	if insertReference != nil {
		return c.JSON(500, "Sorry server error")
	}

	return c.JSON(201, r)
}
