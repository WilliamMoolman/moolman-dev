package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type cycleController struct {
}

// CycleControllerInterface ...
type CycleControllerInterface interface {
	Cycle(c echo.Context) error
	Routes(g *echo.Group)
}

// NewCycleController ...
func NewCycleController() CycleControllerInterface {
	return &cycleController{}
}

// Routes registers route handlers for the health service
func (ctl *cycleController) Routes(g *echo.Group) {
	g.GET("/cycle", ctl.Cycle)
}

// CycleStatus ...
type CycleStatus struct {
	Bikes int64
	Docks int64
}

// StationStatus ...
type StationStatus struct {
	Data struct {
		Stations []Station `json:"stations"`
	} `json:"data"`
}

// Station ...
type Station struct {
	StationID   string `json:"station_id"`
	StationName string
	Bikes       int `json:"num_bikes_available"`
	Docks       int `json:"num_docks_available"`
}

// TODO: Please redo this
// StationList ...
type StationList struct {
	Pollock Station
	Pool    Station
}

// Cycle ...
func (ctl *cycleController) Cycle(c echo.Context) error {
	l := c.Logger()
	l.SetLevel(log.INFO)
	response, err := http.Get("https://gbfs.urbansharing.com/edinburghcyclehire.com/station_status.json")
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	stationData := StationStatus{}
	err = json.Unmarshal(body, &stationData)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	pollock := Station{StationName: "Pollock"}
	pool := Station{StationName: "Royal Commonwealth Pool"}

	for _, station := range stationData.Data.Stations {
		if station.StationID == "264" {
			pollock.Bikes = station.Bikes
			pollock.Docks = station.Docks
		}
		if station.StationID == "246" {
			pool.Bikes = station.Bikes
			pool.Docks = station.Docks
		}
	}

	return c.Render(http.StatusOK, "cycle", StationList{Pollock: pollock, Pool: pool})

}
