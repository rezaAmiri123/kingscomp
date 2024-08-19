package webapp

import (
	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/kingscomp/steps/07_web/internal/webapp/views/pages"
)

func (w *Webapp) lobbyIndex(c echo.Context) error {
	return HTML(c,pages.LobbyPage())
}
