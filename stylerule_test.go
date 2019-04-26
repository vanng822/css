package css

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStyleRuleText(t *testing.T) {
	sr := CSSStyleRule{}
	sr.SelectorText = ".box"
	sr.Styles = make([]*CSSStyleDeclaration, 2)
	sr.Styles[0] = NewCSSStyleDeclaration("width", "10px", 0)
	sr.Styles[1] = NewCSSStyleDeclaration("height", "100px", 0)

	assert.Equal(t, sr.Text(), ".box {\nwidth: 10px;\nheight: 100px\n}")
}
