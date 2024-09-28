package handlers

import (
	"io"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(w io.Writer, c echo.Context, cmp templ.Component) error {
	return cmp.Render(c.Request().Context(), w)
}
