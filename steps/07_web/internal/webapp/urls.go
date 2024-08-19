package webapp

import (
	"context"
	"errors"
	"net/url"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/kingscomp/steps/07_web/internal/entity"
	"github.com/rezaAmiri123/kingscomp/steps/07_web/internal/repository"
	"github.com/rezaAmiri123/kingscomp/steps/07_web/pkg/jsonhelper"
)

func (w *Webapp) urls() {
	lobby := w.e.Group("/lobby")
	lobby.GET("/:lobby_id", w.lobbyIndex)

	auth := w.e.Group("/auth")
	auth.POST("/validate", w.validatInitData, w.authorize)
}

func (w *Webapp) authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		initData := c.Request().Header.Get("Authorization")
		isValid, err := ValidateWebAppInputData(initData)
		if err != nil {
			return err
		}
		if !isValid {
			return c.JSON(403, ResponseError(403, "invalid init data"))
		}
		parsed, _ := url.ParseQuery(initData)
		authTimestamp, _ := strconv.ParseInt(parsed.Get("auth_date"), 10, 64)
		authDate := time.Unix(authTimestamp, 0)
		account := jsonhelper.Decode[entity.Account]([]byte(parsed.Get("user")))

		account, err = w.App.Account.Get(context.Background(), entity.NewID("account", account.ID))
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				return c.JSON(403, ResponseError(403, "account not found"))
			}
			return err
		}
		c.Set("auth_date", authDate)
		c.Set("account", account)
		return next(c)
	}
}
