package webapp

import (
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rezaAmiri123/kingscomp/steps/06_web/internal/service"
)

type Webapp struct {
	App  *service.App
	e    *echo.Echo
	addr string
}

func NewWebApp(app *service.App, addr string) *Webapp {
	e := echo.New()
	wa := &Webapp{
		App:  app,
		e:    e,
		addr: addr,
	}
	wa.urls()
	return wa
}

func (w *Webapp) Start() error {
	w.e.Use(middleware.Recover())
	return w.e.Start(w.addr)
}

func (w *Webapp) StartDev(listener net.Listener) error {
	w.e.Use(middleware.Logger())
	w.e.Use(middleware.Recover())
	return http.Serve(listener, w.e)
}
