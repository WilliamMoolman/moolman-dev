package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/facktoreal/env"
	"github.com/labstack/echo/v4"
)

type telegramController struct {
}

// TelegramControllerInterface ...
type TelegramControllerInterface interface {
	Message(c echo.Context) error
	Routes(g *echo.Group)
}

// NewTelegramController ...
func NewTelegramController() TelegramControllerInterface {
	return &telegramController{}
}

// Routes registers route handlers for the health service
func (ctl *telegramController) Routes(g *echo.Group) {
	g.GET("/message", ctl.Message)
}

type sendMessageReqBody struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

func (ctl *telegramController) Message(c echo.Context) error {
	// Create the request body struct
	reqBody := &sendMessageReqBody{
		ChatID: env.MustGetInt("TELEGRAM_CHAT_ID"),
		Text:   c.QueryParam("message"),
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	// Send a post request with your token
	res, err := http.Post("https://api.telegram.org/bot"+env.MustGetString("TELEGRAM_BOT_API_KEY")+"/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}
