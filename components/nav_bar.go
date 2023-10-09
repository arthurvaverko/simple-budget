package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type NavBar struct {
	app.Compo
	NavItems []NavItem
}

type NavItem struct {
	Icon     string
	Text     string
	Href     string
	IsActive bool
}

func NewNavBar(items []NavItem) *NavBar {
	return &NavBar{
		NavItems: items,
	}
}

func (s *NavBar) Render() app.UI {
	return app.Header().Class("app-header fixed-top").Body(
		app.Div().Class("container-fluid py-2").Body(
			app.Div().Class("row justify-content-between align-items-center").Body(
				app.Div().Class("col-auto").Body(
					app.A().ID("sidepanel-toggler").Class("sidepanel-toggler d-inline-block d-xl-none").Href("/").Body(
						app.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" viewBox="0 0 30 30" role="img"><title>Menu</title><path stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="2" d="M4 7h22M4 15h22M4 23h22"></path></svg>`),
					),
				),
				app.Div().Class("col-auto").Body(),
				app.Div().Class("app-utilities col-auto").Body(
					app.Div().Class("app-utility-item app-user-dropdown dropdown").Body(
						app.A().Class("dropdown-toggle").Href("#").Attr("data-bs-toggle", "dropdown").Body(
							app.Img().Class("app-user-avatar rounded-circle").Src("/web/assets/images/user.png"),
						),
					),
				),
			),
		),

		app.Div().ID("app-sidepanel").Class("app-sidepanel").Body(
			app.Div().ID("sidepanel-drop").Class("sidepanel-drop"),
			app.Div().Class("sidepanel-inner d-flex flex-column").Body(
				app.A().ID("sidepanel-close").Class("sidepanel-close d-xl-none").Href("#").Text("X"),
				app.Div().Class("app-branding").Body(
					app.A().Class("app-logo").Href("/").Body(
						app.Img().Src("/web/assets/app_logo.svg"),
						app.Span().Text("Simple Budget"),
					),
				),
				app.Nav().Class("app-nav app-nav-sidebar").Body(
					app.Ul().Class("app-menu").Body(
						app.Range(s.NavItems).Slice(func(i int) app.UI {
							return app.Li().Class(isActiveClass(s.NavItems[i])).Body(
								app.A().Class("app-menu__item").Href(s.NavItems[i].Href).Body(
									app.Span().Class("app-menu__icon").Body(
										app.Raw(`<i class="bi bi-bar-chart"></i>`),
									),
									app.Span().Class("app-menu__label").Body(
										app.Text(s.NavItems[i].Text),
									),
								),
							)
						}),
					),
				),
			),
		),
	)
}

func isActiveClass(item NavItem) string {
	if item.IsActive {
		return "active"
	}
	return ""
}
