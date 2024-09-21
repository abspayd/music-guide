package handlers

import (
	"net/http"

	"github.com/abspayd/music-guide/pkg/app"
	"github.com/labstack/echo/v4"
)

func GetMain(c echo.Context) error {
	return c.Render(http.StatusOK, "main.html", nil)
}

func GetIntervalCalculator(c echo.Context) error {
	return c.Render(http.StatusOK, "intervals.html", nil)
}

func PostIntervalCalculator(c echo.Context) error {
	note1 := c.Request().PostFormValue("note1")
	note2 := c.Request().PostFormValue("note2")
	interval, err := app.IntervalName(note1, note2)
	if err != nil {
		c.Error(err)
	}
	return c.Render(http.StatusOK, "intervals.html", interval)
}
