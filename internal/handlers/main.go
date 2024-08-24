package handlers

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

/*
func renderBodyTemplate(name string, data interface{}) error {
	return c.Render(http.StatusOK, "main.html", 
}
*/

func HandleMain(c echo.Context) error {
	return c.Render(http.StatusOK, "main.html", nil)
}

func HandleIntervalCalculator(c echo.Context) error {
	return c.Render(http.StatusOK, "intervals.html", nil)
}
