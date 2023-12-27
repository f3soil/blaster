package blaster

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type previewer struct {
	app.Compo

	content string
}

func (x *previewer) OnMount(ctx app.Context) {
	ctx.ObserveState(stateContent).Value(&x.content)
}

func (x *previewer) Render() app.UI {
	return app.Div().ID("previewer").Body(
		app.Strong().Text(x.content))
}
