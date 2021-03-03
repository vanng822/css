package css

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyFrames(t *testing.T) {
	css := Parse(`@keyframes slidein {
			from {
				margin-left: 100%;
				width: 300%;
			}

			to {
				margin-left: 0%;
				width: 100%;
			}
  }`)
	assert.Equal(t, css.CssRuleList[0].Style.Selector.Text(), "slidein")
	assert.Equal(t, css.CssRuleList[0].Type, KEYFRAMES_RULE)
	assert.Equal(t, len(css.CssRuleList[0].Rules), 2)
	assert.Equal(t, css.CssRuleList[0].Rules[0].Style.Selector.Text(), "from")
	assert.Equal(t, css.CssRuleList[0].Rules[0].Style.Styles[0].Value.Text(), "100%")
	assert.Equal(t, css.CssRuleList[0].Rules[0].Style.Styles[1].Value.Text(), "300%")
	assert.Equal(t, css.CssRuleList[0].Rules[1].Style.Selector.Text(), "to")
}

func TestKeyFramesPercent(t *testing.T) {
	css := Parse(`@keyframes identifier {
		0% { top: 0; }
		50% { top: 30px; left: 20px; }
		50% { top: 10px; }
		100% { top: 0; }
  }`)
	assert.Equal(t, css.CssRuleList[0].Style.Selector.Text(), "identifier")
	assert.Equal(t, css.CssRuleList[0].Type, KEYFRAMES_RULE)
	assert.Equal(t, len(css.CssRuleList[0].Rules), 4)
	assert.Equal(t, css.CssRuleList[0].Rules[0].Style.Selector.Text(), "0%")
	assert.Equal(t, css.CssRuleList[0].Rules[0].Style.Styles[0].Value.Text(), "0")
	assert.Equal(t, css.CssRuleList[0].Rules[1].Style.Selector.Text(), "50%")
	assert.Equal(t, css.CssRuleList[0].Rules[2].Style.Selector.Text(), "50%")
	assert.Equal(t, css.CssRuleList[0].Rules[3].Style.Selector.Text(), "100%")
}

func TestWebKitKeyFramesPercent(t *testing.T) {
	css := Parse(`@-webkit-keyframes identifier {
		0% { top: 0; }
		50% { top: 30px; left: 20px; }
		50% { top: 10px; }
		100% { top: 0; }
  }`)
	assert.Equal(t, css.CssRuleList[0].Style.Selector.Text(), "identifier")
	assert.Equal(t, css.CssRuleList[0].Type, WEBKIT_KEYFRAMES_RULE)
	assert.Equal(t, len(css.CssRuleList[0].Rules), 4)
	assert.Equal(t, css.CssRuleList[0].Rules[0].Style.Selector.Text(), "0%")
	assert.Equal(t, css.CssRuleList[0].Rules[0].Style.Styles[0].Value.Text(), "0")
	assert.Equal(t, css.CssRuleList[0].Rules[1].Style.Selector.Text(), "50%")
	assert.Equal(t, css.CssRuleList[0].Rules[2].Style.Selector.Text(), "50%")
	assert.Equal(t, css.CssRuleList[0].Rules[3].Style.Selector.Text(), "100%")
}
