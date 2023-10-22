package app

import (
	"github.com/arthurvaverko/simple-budget/app/components"
	"github.com/arthurvaverko/simple-budget/app/pages"
	"github.com/labstack/echo/v4"
)

var appPages = []components.AppPage{
	pages.Index(),
}

func RegisterRoutes(e *echo.Echo) {
	for _, page := range appPages {
		e.GET(page.Route(), func(c echo.Context) error {
			res, err := page.Render()
			if err != nil {
				c.Logger().Error(err)
				return err
			}
			return c.HTMLBlob(200, res.Bytes())
		})
	}
}
