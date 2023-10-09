package components

//
//import (
//	"github.com/maxence-charriere/go-app/v9/pkg/app"
//)
//
//type NavBar struct {
//	app.Compo
//	NavItems []NavItem
//}
//
//type NavItem struct {
//	Icon     string
//	Text     string
//	Href     string
//	IsActive bool
//}
//
//func NewNavBar(items []NavItem) *NavBar {
//	return &NavBar{
//		NavItems: items,
//	}
//}
//
//func (s *NavBar) Render() app.UI {
//	return app.Nav().
//		ID("main-van").
//		Class("navbar bg-body-tertiary fixed-top").
//		Body(
//			app.Div().Class("container-fluid").Body(
//				app.A().Class("navbar-brand").Href("/").Body(
//					app.Img().Src("/web/favicon.png").Class("d-inline-block align-top").Style("height", "30px"),
//					app.Span().Class("ms-2").Text("Simple Budget"),
//				),
//				app.Button().Class("navbar-toggler").
//					Attr("data-bs-toggle", "offcanvas").
//					Attr("data-bs-target", "#offcanvasNavbar").
//					Aria("controls", "offcanvasNavbar").
//					Aria("expanded", "false").
//					Aria("label", "Toggle navigation").Body(
//					app.Span().Class("navbar-toggler-icon"),
//				),
//
//				app.Div().Class("offcanvas offcanvas-end").
//					TabIndex(-1).
//					ID("offcanvasNavbar").
//					Aria("aria-labelledby", "offcanvasNavbarLabel").
//					Body(
//						app.Div().Class("offcanvas-header").Body(
//							app.H5().Class("offcanvas-title").ID("offcanvasNavbarLabel").Body(
//								app.Img().Src("/web/favicon.png").Class("d-inline-block align-top").Style("height", "30px"),
//								app.Span().Text("Simple Budget"),
//							),
//							app.Button().Class("btn-close text-reset").Attr("data-bs-dismiss", "offcanvas").Aria("label", "Close"),
//						),
//						app.Div().Class("offcanvas-body").Body(
//							app.Ul().Class("navbar-nav justify-Content-end flex-grow-1 pe-3").Body(
//								app.Range(s.NavItems).Slice(func(i int) app.UI {
//									return app.Li().Class("nav-item").Body(
//										app.A().Href(s.NavItems[i].Href).Class(isActiveClass(s.NavItems[i])).Class("nav-link icon-link").Body(
//											app.I().Style("height", "auto").Class(s.NavItems[i].Icon),
//											app.Span().Text(s.NavItems[i].Text),
//										),
//									)
//								}),
//							),
//						),
//					),
//			),
//		)
//}
//
//func isActiveClass(item NavItem) string {
//	if item.IsActive {
//		return "active"
//	}
//	return ""
//}
