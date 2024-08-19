package webapp

import (
	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/kingscomp/steps/07_web/internal/entity"
)

type validateInitDataRequest struct {
	InitData string `json:"initData"`
}

func (w *Webapp) validatInitData(c echo.Context) error {
	acc := c.Get("account").(entity.Account)
	return c.JSON(200, ResponseOk(200, J{
		"is_valid": true,
		"account":  acc,
	}))
}
