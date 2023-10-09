package main

import (
	"github.com/arthurvaverko/simple-budget/pages"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"log"
	"net/http"
)

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &pages.Dashboard{})
	app.Route("/report-expense", &pages.ReportExpense{})
	app.Route("/recurring-expenses", &pages.ReportExpenses{})

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	log.Println("starting...")
	http.Handle("/", &app.Handler{
		Name:        "Simple Budget",
		Title:       "Simple Budget",
		Author:      "Arthur Vaverko",
		Description: "Just so you know how much is left...",
		Styles: []string{
			"/web/assets/css/portal.css",
		},
		Scripts: []string{
			//https://themes.3rdwavemedia.com/bootstrap-templates/startup/portal-free-bootstrap-admin-dashboard-template-for-developers/
			"web/assets/js/app.js",
			"/web/assets/plugins/fontawesome/js/all.min.js",
		},
		HTML: func() app.HTMLHtml {
			return app.Html().Attr("lang", "en")
		},
		Icon: app.Icon{
			SVG: "/web/assets/app_logo.svg",
		},
		Image: "/web/favicon.png",
		// TODO: clear cache only while dev
		RawHeaders: []string{
			`<script>
				caches.keys().then((keyList) => Promise.all(keyList.map((key) => caches.delete(key))));
				console.log("cleared cache");
			</script>`,
		},
	})

	log.Println("listening on http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
