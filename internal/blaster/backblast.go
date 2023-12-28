package blaster

import (
	"fmt"
	"strings"

	"cloud.google.com/go/civil"
)

type Backblast struct {
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

func (x *Backblast) Markdown() string {
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
