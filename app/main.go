package main

import (
	"github.com/facktoreal/env"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	sentryEcho "github.com/getsentry/sentry-go/echo"
)

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
	e.File("/", "public/index.html")

	// Once it's done, you can attach the handler as one of your middleware
	e.Use(sentryEcho.New(sentryEcho.Options{Repanic: true}))

	// Services
	var ()

	// Core endpoints
	//controllers.NewHealthController(db).Routes(e.Group("api"))

	e.Logger.Infof("Server started, v%s | port: %s", echo.Version, port)

	e.Logger.Fatal(e.Start(":" + port))
}
