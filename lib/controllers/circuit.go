package controllers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/html"
)

type circuitController struct {
}

// CircuitControllerInterface ...
type CircuitControllerInterface interface {
	Laundry(c echo.Context) error
	Routes(g *echo.Group)
}

// NewCircuitController ...
func NewCircuitController() CircuitControllerInterface {
	return &circuitController{}
}

// Routes registers route handlers for the health service
func (ctl *circuitController) Routes(g *echo.Group) {
	g.GET("/laundry", ctl.Laundry)
}

// LaundryStatus ...
type LaundryStatus struct {
	Washing string
	Drying  string
}

// Laundry ...
func (ctl *circuitController) Laundry(c echo.Context) error {
	l := c.Logger()
	l.SetLevel(log.INFO)
	response, err := http.Get("https://www.circuit.co.uk/circuit-view/laundry-site/?site=5669")
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	defer response.Body.Close()
	z := html.NewTokenizer(response.Body)
	end := false
	washers := "n/a"
	dryers := "n/a"
	for end == false {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			end = true
		case tt == html.StartTagToken:
			t := z.Token()

			isP := t.Data == "p"
			if isP {
				z.Next()
				a := z.Token()
				if strings.HasSuffix(a.Data, "Washers Available") {
					washers = a.Data
				}
				if strings.HasSuffix(a.Data, "Dryers Available") {
					dryers = a.Data
				}
			}
		}
	}
	return c.Render(http.StatusOK, "laundry", LaundryStatus{
		Washing: washers,
		Drying:  dryers,
	})

}
