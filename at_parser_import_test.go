package css

import (
	"testing"

	"github.com/gorilla/css/scanner"
	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	css := Parse(`@import url("fineprint.css") print;
					@import url("bluish.css") projection, tv;
					@import 'custom.css';
					@import url("chrome://communicator/skin/");
					@import "common.css" screen, projection;
					@import url('landscape.css') screen and (orientation:landscape);`)

	assert.Equal(t, css.CssRuleList[0].Style.Selector.Text(), "url(\"fineprint.css\") print")
	assert.Equal(t, css.CssRuleList[1].Style.Selector.Text(), "url(\"bluish.css\") projection, tv")
	assert.Equal(t, css.CssRuleList[2].Style.Selector.Text(), "'custom.css'")
	assert.Equal(t, css.CssRuleList[3].Style.Selector.Text(), "url(\"chrome://communicator/skin/\")")
	assert.Equal(t, css.CssRuleList[4].Style.Selector.Text(), "\"common.css\" screen, projection")
	assert.Equal(t, css.CssRuleList[5].Style.Selector.Text(), "url('landscape.css') screen and (orientation:landscape)")

	assert.Equal(t, css.CssRuleList[0].Type, IMPORT_RULE)
	assert.Equal(t, css.CssRuleList[1].Type, IMPORT_RULE)
	assert.Equal(t, css.CssRuleList[2].Type, IMPORT_RULE)
	assert.Equal(t, css.CssRuleList[3].Type, IMPORT_RULE)
	assert.Equal(t, css.CssRuleList[4].Type, IMPORT_RULE)
	assert.Equal(t, css.CssRuleList[5].Type, IMPORT_RULE)
}

func TestImportIgnore(t *testing.T) {
	css := parseAtNoBody(scanner.New(` url("fineprint.css") print`), IMPORT_RULE)
	assert.Nil(t, css)
}
