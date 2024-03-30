package webapp

import (
	"github.com/labstack/echo/v4"
	"github.com/a-h/templ"
)

type J map[string]any

func HTML(c echo.Context, cmp templ.Com)