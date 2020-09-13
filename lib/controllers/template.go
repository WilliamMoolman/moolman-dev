package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type templateController struct {
}

// TemplateControllerInterface ...
type TemplateControllerInterface interface {
	Hello(c echo.Context) error
	Routes(g *echo.Group)
}

// NewTemplateController ...
func NewTemplateController() TemplateControllerInterface {
	return &templateController{}
}

// Routes registers route handlers for the health service
func (ctl *templateController) Routes(g *echo.Group) {
	g.GET("/templates/hello", ctl.Hello)
}

// Hello ...
func (ctl *templateController) Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello.html", map[string]interface{}{
		"name": "Stranger!",
	})

}
