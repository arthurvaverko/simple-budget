package main

import (
	"embed"
	"fmt"
	"github.com/arthurvaverko/simple-budget/app"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"net/http"
)

//go:embed public/*
var publicFs embed.FS

var loggerFormat = "${level} ${time_rfc3339} ${id} ${method} ${uri} ${status} ${error} ${latency_human} \n"

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: loggerFormat}))
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

	fmt.Println("Starting server on http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
