package css

import (
	"testing"

	"github.com/gorilla/css/scanner"
	"github.com/stretchr/testify/assert"
)

func TestCharsetDoubleQ(t *testing.T) {
	css := Parse(`@charset "UTF-8";`)

	assert.Equal(t, css.CssRuleList[0].Style.Selector.Text(), "\"UTF-8\"")
	assert.Equal(t, css.CssRuleList[0].Type, CHARSET_RULE)
}

func TestCharsetSingleQ(t *testing.T) {
	css := Parse(`@charset 'iso-8859-15';`)

	assert.Equal(t, css.CssRuleList[0].Style.Selector.Text(), "'iso-8859-15'")
	assert.Equal(t, css.CssRuleList[0].Type, CHARSET_RULE)
}

func TestCharsetIgnore(t *testing.T) {
	css := parseAtNoBody(scanner.New(` 'iso-8859-15'`), CHARSET_RULE)

	assert.Nil(t, css)
}
