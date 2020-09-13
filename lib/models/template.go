package models

import (
	"io"
	"text/template"

	"github.com/labstack/echo"
)

// TemplateRenderer ...
type TemplateRenderer struct {
	Templates *template.Template
}

// Render ...
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}
