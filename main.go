package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"

	"github.com/f3soil/blaster/internal/blaster"
)

func main() {
	app.Route("/", &blaster.App{})
	app.RunWhenOnBrowser()
	http.Handle("/", &app.Handler{
		Name:        "Blaster",
		Description: "Blaster helps you Backblast",
		Resources:   app.GitHubPages("REPOSITORY_NAME"),
		Styles: []string{
			"/web/app.css",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
