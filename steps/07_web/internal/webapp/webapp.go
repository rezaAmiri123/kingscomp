package webapp

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rezaAmiri123/kingscomp/steps/07_web/internal/service"
	"github.com/sirupsen/logrus"
)

//go:embed static
var embededfiles embed.FS

func getFileSystem() http.FileSystem {
	fSys, err := fs.Sub(embededfiles, "static")
	if err != nil {
		logrus.WithError(err).Panicln("couldn't init static embedding")
	}
	return http.FS(fSys)
}

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
	wa.static()
	return wa
}

func (w *Webapp) Start() error {
	w.e.Use(middleware.Recover())
	return w.e.Start(w.addr)
}

func (w *Webapp) Shutdown(ctx context.Context) error {
	return w.e.Shutdown(ctx)
}

func (w *Webapp) StartDev(listener net.Listener) error {
	w.e.Use(middleware.Logger())
	w.e.Use(middleware.Recover())
	return http.Serve(listener, w.e)
}

func (w *Webapp) static() {
	assestHandler := http.FileServer(getFileSystem())
	w.e.GET("/static/*",
		echo.WrapHandler(http.StripPrefix("/static/", assestHandler)),
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				c.Response().Header().Set(
					"Cache-Control",
					fmt.Sprintf("public,max-age=%d", int((time.Hour*24).Seconds())),
				)
				err := next(c)
				if err != nil {
					return err
				}
				return nil
			}
		},
	)
}
