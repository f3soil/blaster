package main

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"

	"github.com/f3soil/blaster/internal/blaster"
)

func main() {
	app.Route("/", &blaster.App{})
	app.RunWhenOnBrowser()
	err := app.GenerateStaticWebsite("docs", &app.Handler{
		Name:        "Blaster",
		Description: "Blaster helps you Backblast",
		Resources:   app.GitHubPages("blaster"),
		Scripts: []string{
			"/web/blaster.js",
		},
		Styles: []string{
			"/web/blaster.css",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
