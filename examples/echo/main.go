package main

import (
	"context"
	"fmt"
	"net/http"

	log "github.com/kunstack/klog"
	"github.com/labstack/echo/v4"
)

// Do some operations...
func someService(ctx context.Context) {
	l := log.FromContext(ctx)
	defer l.Flush()
	l.Warningln("this is someService...")
}

// curl -H "X-Request-ID: 0F0623A4-0980-47FB-8257-664FA5761E6C" http://localhost:8080/ping

func main() {
	app := echo.New()
	app.Use(
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(ctx echo.Context) error {
				newCtx := log.WithContext(
					ctx.Request().Context(),
					log.StrField("request-id", ctx.Request().Header.Get("x-request-id")),
				)
				ctx.SetRequest(ctx.Request().WithContext(newCtx))
				return next(ctx)
			}
		},
	)

	app.GET("/ping", func(ctx echo.Context) error {
		l := log.FromContext(ctx.Request().Context())
		defer l.Flush()
		l.Infoln("this is test message")
		someService(ctx.Request().Context())
		return ctx.String(http.StatusOK, fmt.Sprintf("reuest-id is %s", l.Fields().Get("request-id")))
	})

	if err := app.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
