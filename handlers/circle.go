package handlers

import (
	"github.com/abspayd/music-guide/views/circle"
	"github.com/labstack/echo/v4"
)

func GetCircleOfFifths(c echo.Context) error {
	return Render(c.Response().Writer, c, circle.CircleOfFifths())
}
