package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMardown(t *testing.T) {
	assert.Exactly(t, `## Release Note

### Features

* 7702903 feat: lorem ipsum dolor sit amet
* 315h3dk feat(foo): cras ultricies ligula sed magna dictum porta

### Fixes

* 70f90c4 fix(foo): donec sollicitudin molestie malesuada

### Documentation

* 4272900 docs: aliquam quis turpis eget elit sodales scelerisque
* 23f2dg4 docs(foo): sed porttitor lectus nibh

`,
		NewDocument().With(
			H2("Release Note"),
			H3("Features"),
			UL().With(
				LI("7702903 feat: lorem ipsum dolor sit amet"),
				LI("315h3dk feat(foo): cras ultricies ligula sed magna dictum porta"),
			),
			H3("Fixes"),
			UL().With(
				LI("70f90c4 fix(foo): donec sollicitudin molestie malesuada"),
			),
			H3("Documentation"),
			UL().With(
				LI("4272900 docs: aliquam quis turpis eget elit sodales scelerisque"),
			).With(
				LI("23f2dg4 docs(foo): sed porttitor lectus nibh"),
			),
		).String(),
	)

	assert.Exactly(t, `## Release Note

* 70f90c4 fix: foo
* 7702903 feat: bar

`,
		NewDocument().With(
			H2("Release Note"),
			UL().With(
				LI("70f90c4 fix: foo"),
				LI("7702903 feat: bar"),
			),
		).String(),
	)
}
