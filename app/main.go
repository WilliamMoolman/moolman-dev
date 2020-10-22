package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/WilliamMoolman/moolman-dev/lib/controllers"
	"github.com/facktoreal/env"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer ...
type TemplateRenderer struct {
	templates *template.Template
}

// Render ...
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// Hide banner
	e.HideBanner = true

	// Load environment vars
	if err := env.Init(); err != nil {
		e.Logger.Fatalf("Unable to load environment variables, err: %s", err.Error())
	}

	port := env.MustGetString("PORT")
	if port == "" {
		port = "8080"
		e.Logger.Infof("Defaulting to port %s", port)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	// Templates
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = renderer

	// Core endpoints
	controllers.NewTemplateController().Routes(e.Group(""))
	controllers.NewCircuitController().Routes(e.Group(""))
	controllers.NewCycleController().Routes(e.Group(""))
	controllers.NewTelegramController().Routes(e.Group("api"))
	fs := http.FileServer(http.Dir("public"))
	e.GET("/public/*", echo.WrapHandler(http.StripPrefix("/public/", fs)))
	e.File("/", "public/index.html")

	e.Logger.Infof("Server started, v%s | port: %s", echo.Version, port)

	e.Logger.Fatal(e.Start(":" + port))
}
