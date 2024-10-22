package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"strings"

	"github.com/abspayd/music-guide/pkg/app"
	"github.com/abspayd/music-guide/views/components"
	"github.com/abspayd/music-guide/views/intervals"
)

func GetIntervalCalculator(c echo.Context) error {
	return Render(c.Response().Writer, c, intervals.Intervals(nil))
}

func PostIntervalCalculator(c echo.Context) error {
	note1 := c.Request().PostFormValue("note1")
	note2 := c.Request().PostFormValue("note2")
	interval, err := app.IntervalName(note1, note2)

	if err != nil {
		// One or both of the inputs are invalid so just leave without swapping
		c.Response().Header().Set("HX-Reswap", "none")
		return nil
	}

	// TODO: just a start to better output (I want it to show the inputs that lead to the answer)
	// output := fmt.Sprintf("%s -> %s = %s", note1, note2, interval)
	// return Render(c.Response().Writer, c, intervals.IntervalEntry(output))

	return Render(c.Response().Writer, c, intervals.IntervalEntry(interval))
}

func ValidateNote(c echo.Context) error {
	// get the input field(s?) that triggered validation
	header := c.Request().Header
	_, ok := header["Hx-Request"]
	if !ok {
		log.Println("Non-HTMX validation recieved. Not sure how this can happen.")
		return nil
	}

	hx_trigger := header["Hx-Trigger"]

	id := strings.Join(hx_trigger, "")
	value := c.Request().FormValue(id)
	label := c.Request().FormValue("label")

	_, err := app.NewPitch(value)
	if err == nil {
		return Render(c.Response().Writer, c, components.TextInput(label, id, value))
	}

	// determine if the input was wrong or just empty
	error_msg := ""
	if len(value) == 0 {
		error_msg = "This field is required"
	} else {
		error_msg = "Invalid pitch"
	}

	return Render(c.Response().Writer, c, components.TextInputInvalid(label, id, value, error_msg))
}
