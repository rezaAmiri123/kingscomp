package webapp

import (
	"embed"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/kingscomp/internal/gameserver"
	"github.com/rezaAmiri123/kingscomp/internal/service"
	"gopkg.in/telebot.v3"
)

//go:embed static
var embededFiles embed.FS

type WebApp struct{
	App *service.App
	e *echo.Echo
	addr string
	gs *gameserver.GameServer
	bot *telebot.Bot
}

func NewWebApp(
	app *service.App,
	addr string,
	gs *gameserver.GameServer,
	bot *telebot.Bot,
)*WebApp{
	e := echo.New()
	wa := &WebApp{
		App: app,
		e: e,
		addr: addr,
		bot: bot,
	}
	wa.urls()
	wa.static()
	return wa
}

func(w *WebApp)Start()error{
	w.e.Use(middleware.Recover())
}