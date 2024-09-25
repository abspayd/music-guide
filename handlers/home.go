package handlers

import (
	"github.com/abspayd/music-guide/views/templates"
	"github.com/labstack/echo/v4"
)

func GetHome(c echo.Context) error {
	return Render(c.Response().Writer, c, templates.Home())
}
