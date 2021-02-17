package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/kunstack/klog"
)

// Do some operations...
func someService(ctx context.Context) {
	l := log.FromContext(ctx)
	defer l.Flush()
	l.Warningln("this is someService...")
}

func main() {
	router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		newCtx := log.WithContext(
			ctx.Request.Context(),
			log.StrField("request-id", ctx.Request.Header.Get("x-request-id")),
		)
		ctx.Request = ctx.Request.WithContext(newCtx)
	})

	router.GET("/ping", func(ctx *gin.Context) {
		l := log.FromContext(ctx.Request.Context())
		defer l.Flush()
		l.Infoln("this is test message")
		someService(ctx.Request.Context())
		ctx.String(http.StatusOK, fmt.Sprintf("reuest-id is %s", l.Fields().Get("request-id")))
	})

	router.Run(":8080")
}
