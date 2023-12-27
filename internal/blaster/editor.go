package blaster

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type editor struct {
	app.Compo
}

func (x *editor) Render() app.UI {
	return app.Div().ID("editor").Body(
		app.Input().
			OnKeyPress(x.onKeyPress).
			Type("textbox").
			AutoFocus(true),
	)
}

func (x *editor) onKeyPress(ctx app.Context, e app.Event) {
	ctx.SetState(stateContent, e.String())
}
