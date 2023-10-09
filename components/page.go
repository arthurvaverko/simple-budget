package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Page struct {
	app.Compo
	ContentElements []app.UI
}

func NewPage() *Page {
	return &Page{}
}

func (p *Page) Content(v ...app.UI) *Page {
	p.ContentElements = v
	return p
}

func (p *Page) Render() app.UI {
	navItems := []NavItem{
		{"bi bi-bar-chart", "Status Dashboard", "/", true},
		{"bi bi-input-cursor-text", "Report Expense", "/report-expense", false},
		{"bi bi-recycle", "Recurring Expenses", "/recurring-expenses", false},
	}

	for i, _ := range navItems {
		navItems[i].IsActive = app.Window().URL().Path == navItems[i].Href
	}

	return app.Div().
		Body(
			NewNavBar(navItems),
			app.Div().Class("app-content pt-3 p-md-3 p-lg-4").Body(
				app.Div().Class("container-xl").Body(
					p.ContentElements...,
				),
			),
		)
}
