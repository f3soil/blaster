package blaster

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	stateContent = "content"
)

type App struct {
	app.Compo

	name string
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (x *App) Render() app.UI {
	return app.Div().ID("blaster").Body(
		&editor{},
		&previewer{},
	)
}
