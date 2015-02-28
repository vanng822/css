package css

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImport(t *testing.T) {
	css := Parse(`@import url("fineprint.css") print;
					@import url("bluish.css") projection, tv;
					@import 'custom.css';
					@import url("chrome://communicator/skin/");
					@import "common.css" screen, projection;
					@import url('landscape.css') screen and (orientation:landscape);`)

	assert.Equal(t, css.CssRuleList[0].Style.SelectorText, "url(\"fineprint.css\") print")
	assert.Equal(t, css.CssRuleList[1].Style.SelectorText, "url(\"bluish.css\") projection, tv")
	assert.Equal(t, css.CssRuleList[2].Style.SelectorText, "'custom.css'")
	assert.Equal(t, css.CssRuleList[3].Style.SelectorText, "url(\"chrome://communicator/skin/\")")
	assert.Equal(t, css.CssRuleList[4].Style.SelectorText, "\"common.css\" screen, projection")
	assert.Equal(t, css.CssRuleList[5].Style.SelectorText, "url('landscape.css') screen and (orientation:landscape)")
}

