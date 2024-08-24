package main

import (
	"html/template"
	"io"

	"github.com/abspayd/music-guide/internal/routes"
	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Static("/static", "web/static")

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("./web/templates/*.html")),
	}
	e.Renderer = renderer

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
