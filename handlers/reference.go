package handlers

import (
	"shortener-service/helpers" // Load helpers
	"shortener-service/models"

	"github.com/labstack/echo"
)

// GetAllUrl list of all url reference
func GetAllUrl(c echo.Context) error {
	response := &helpers.Response{&models.Reference{2, "https://www.google.com", "https://kw.com/x323"}, "All URL"}

	return c.JSON(200, response)
}
