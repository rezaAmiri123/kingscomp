package webapp

import (
	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/kingscomp/steps/08_question/internal/webapp/views/pages"
)

func (w *Webapp) lobbyIndex(c echo.Context) error {
	return HTML(c, pages.LobbyPage(c.Param("lobbyId")))
}

func (w *Webapp) lobbyReady(c echo.Context) error {
	account := getAccount(c)
	lobby := getLobby(c)

	if err := w.App.Lobby.UpdateUserState(c.Request().Context(),
		lobby.ID, account.ID, "isReady", true); err != nil {
		return err
	}
	return c.JSON(200, ResponseOk(200, "done"))
}
