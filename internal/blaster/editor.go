package blaster

import (
	"fmt"
	"strings"

	"cloud.google.com/go/civil"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type editor struct {
	app.Compo

	When civil.Date
	AO   string
	What string

	Total int

	Conditions string

	WarmUp        string
	TheThang      string
	Mary          string
	Announcements string
	CoT           string
}

func (x *editor) Render() app.UI {
	return app.Div().ID("editor").Body(
		app.Form().Body(
			app.Label().Text("When"),
			app.Input().ID("when").
				OnChange(x.valueTo(&x.When)),
			app.Br(),
			app.Label().Text("AO"),
			app.Select().ID("ao").
				Body(
					app.Option().Text("The Mine"),
				).
				OnChange(x.valueTo(&x.AO)),
			app.Br(),
			app.Label().Text("What"),
			app.Input().ID("what").
				OnChange(x.valueTo(&x.What)),
			app.Br(),
			app.Label().Text("Total"),
			app.Input().ID("total").
				OnChange(x.valueTo(&x.Total)),
			app.Br(),
			app.Label().Text("Welcome"),
			app.Textarea().ID("welcome"),
			app.Br(),
			app.Label().Text("Warm-Up"),
			app.Textarea().ID("warm-up").
				Attr("rows", 4).
				Attr("cols", 60).
				OnChange(x.valueTo(&x.WarmUp)),
			app.Br(),
			app.Label().Text("The Thang"),
			app.Textarea().ID("the-thang").
				Attr("rows", 4).
				Attr("cols", 60).
				OnChange(x.valueTo(&x.TheThang)),
			app.Br(),
			app.Label().Text("Mary"),
			app.Textarea().ID("mary").
				Attr("rows", 4).
				Attr("cols", 60).
				OnChange(x.valueTo(&x.Mary)),
			app.Br(),
			app.Label().Text("Announcements"),
			app.Textarea().ID("announcements").
				Attr("rows", 4).
				Attr("cols", 60).
				OnChange(x.valueTo(&x.Announcements)),
			app.Br(),
			app.Label().Text("CoT"),
			app.Textarea().ID("cot").
				Attr("rows", 4).
				Attr("cols", 60).
				OnChange(x.valueTo(&x.CoT)),
		),
	)
}

func (x *editor) valueTo(v any) func(ctx app.Context, e app.Event) {
	return func(ctx app.Context, e app.Event) {
		// TODO: Need to store the value into state? Maybe JSON marshal the whole Backblast struct? app.Persist?
		x.ValueTo(v)(ctx, e)
		ctx.SetState(stateContent, x.Markdown())
	}
}

func (x *editor) Markdown() string {
	lines := []string{
		"*BACKBLAST*",
		"",
		"*When:* " + x.When.String(),
		"*AO:* " + strings.TrimSpace(x.AO),
		"*What:* " + strings.TrimSpace(x.What),
		"",
		"*PAX:* ???",
		"*FNGs:* ? - ???",
		"*Q:* ???",
		fmt.Sprintf("*Total:* %d", x.Total),
		"",
		"*Conditions:* " + strings.TrimSpace(x.Conditions),
		"",
		"*Welcome*",
		"",
		"* *DiCCS:* *Disclaimer*, *CPR*, *Cell* Phones, *Safety*",
		"* *FOOPC:* Always *Free*, *Open* to all men, held *Outdoors* rain or shine, *Peer-Led* in a rotating fashion, always ends in a *CoT*",
		"* Mission: To plant, grow, and serve small workout groups for men, for the invigoration of male community leadership.",
		"",
		"*Warm-Up*",
		"",
		strings.TrimSpace(x.WarmUp),
		"",
		"*The Thang*",
		"",
		strings.TrimSpace(x.TheThang),
		"",
		"*Mary*",
		"",
		strings.TrimSpace(x.Mary),
		"",
		"*CoR + NoR*",
		"",
		"*Announcements*",
		"",
		strings.TrimSpace(x.Announcements),
		"",
		"*CoT*",
		"",
		strings.TrimSpace(x.CoT),
	}
	return strings.Join(lines, "\n")
}
