package releasenote

import (
	"bytes"
	"text/template"

	"releaseros/internal/config"
)

type Footer struct {
	LatestTag   string
	PreviousTag string
}

func (footer Footer) Generate(config config.Config) (string, error) {
	var out bytes.Buffer
	footerText := config.Footer

	t, err := template.New("footer").Funcs(FunctionsForTemplate).Parse(footerText)
	if err != nil {
		return "", err
	}
	err = t.Execute(&out, footer)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
