package releasenote

import (
	"bytes"
	"text/template"

	"releaseros/internal/context"
)

type Footer struct {
	LatestTag   string
	PreviousTag string
}

func (footer Footer) Generate(ctx *context.Context) (string, error) {
	var out bytes.Buffer
	footerText := ctx.Config.Footer

	t, err := template.New("footer").Parse(footerText)
	if err != nil {
		return "", err
	}
	err = t.Execute(&out, footer)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
