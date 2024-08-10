package main

import (
	"io"
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/abspayd/music-guide/router"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, c)
}

func main() {
	e := echo.New()

	renderer := &Template {
		templates: template.Must(template.ParseGlob("*.html")),
	}

	e.Renderer = renderer

	router.HandleRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))

	// -- deprecated --
	// port := 3000
	// fmt.Printf("Starting server on port %d.\n", port)
	// router.HandleRoutes()
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
