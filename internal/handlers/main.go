package handlers

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func HandleMain(c echo.Context) error {
	return c.Render(http.StatusOK, "main.html", nil)
}
