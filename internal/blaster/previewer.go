package blaster

import (
	"fmt"

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
	previewerTextID := "previewer-text"
	return app.Div().ID("previewer").Body(
		app.Textarea().ID(previewerTextID).
			Attr("rows", 40).
			Attr("cols", 60).
			ReadOnly(true).
			Text(x.content),
		app.Br(),
		app.Button().Attr("onclick", fmt.Sprintf("copyToClipboard(%q)", previewerTextID)).Text("Copy"),
	)
}
