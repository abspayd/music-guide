package handlers

import (

	"github.com/abspayd/music-guide/pkg/app"
	"github.com/abspayd/music-guide/views/intervals"

	"github.com/labstack/echo/v4"
)

func GetIntervalCalculator(c echo.Context) error {
	// Read cookies here, not on every POST
	history := []string{}
	// cookie, err := c.Cookie("interval_history")
	// if err != nil {
	// 	log.Println(err)
	// }

	return Render(c.Response().Writer, c, intervals.Intervals(history))
}

func PostIntervalCalculator(c echo.Context) error {
	note1 := c.Request().PostFormValue("note1")
	note2 := c.Request().PostFormValue("note2")
	interval, err := app.IntervalName(note1, note2)
	if err != nil {
		c.Error(err)
	}

	return Render(c.Response().Writer, c, intervals.IntervalEntry(interval))
}
