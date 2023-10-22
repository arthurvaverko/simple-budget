package main

import (
	"embed"
	"github.com/arthurvaverko/simple-budget/app"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"golang.org/x/time/rate"
	"net/http"
)

//go:embed public/*
var publicFs embed.FS

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	app.RegisterRoutes(e)

	e.GET("/health", func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json")
		return c.String(http.StatusOK, `{"status": "ok"}`)
	})

	var publicStaticContentHandler = echo.WrapHandler(http.FileServer(http.FS(publicFs)))
	var publicContentRewriteRule = middleware.Rewrite(map[string]string{"/*": "/public/$1"})
	e.GET("/*", publicStaticContentHandler, publicContentRewriteRule)

	e.Logger.Info("Starting server on http://localhost:4040")
	e.Logger.Fatal(e.Start(":4040"))
}
