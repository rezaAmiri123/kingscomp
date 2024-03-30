package webapp

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/kingscomp/internal/config"
	"gopkg.in/telebot.v3"
)

func(w *WebApp)webhook(c echo.Context)error{
	if c.Param("token")!=config.Default.BotToken{
		return c.String(403, "bad api token")
	}

	update := new(telebot.Update)
	if err := c.Bind(update);err!= nil{
		return err
	}

	w.bot.ProcessUpdate(*update)
	return c.String(http.StatusOK, "OK")
}
