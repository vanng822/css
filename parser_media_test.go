package css

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedia(t *testing.T) {
	css := Parse(`@media only screen and (max-width: 600px) {
			    table[class="body"] img {
			        width: auto !important;
			        height: auto !important
			        }
			    table[class="body"] center {
			        min-width: 0 !important
			        }
			    table[class="body"] .container {
			        width: 95% !important
			        }
			    table[class="body"] .row {
			        width: 100% !important;
			        display: block !important
			        }
        		}`)

	//fmt.Println(css)
	assert.Equal(t, css.CssRuleList[0].Style.Selector.Text(), "only screen and (max-width: 600px)")
	assert.Equal(t, css.CssRuleList[0].Type, MEDIA_RULE)
	assert.Equal(t, len(css.CssRuleList[0].Rules), 4)
	assert.Equal(t, css.CssRuleList[0].Rules[0].Style.Selector.Text(), "table[class=\"body\"] img")
	assert.Equal(t, css.CssRuleList[0].Rules[0].Style.Styles[0].Value.Text(), "auto")
	assert.Equal(t, css.CssRuleList[0].Rules[0].Style.Styles[1].Important, true)
	assert.Equal(t, css.CssRuleList[0].Rules[1].Style.Selector.Text(), "table[class=\"body\"] center")
	assert.Equal(t, css.CssRuleList[0].Rules[2].Style.Selector.Text(), "table[class=\"body\"] .container")
	assert.Equal(t, css.CssRuleList[0].Rules[3].Style.Selector.Text(), "table[class=\"body\"] .row")

}

func TestMediaMulti(t *testing.T) {
	css := Parse(`
				table.one {
					width: 30px;
				}
				@media only screen and (max-width: 600px) {
				    table[class="body"] img {
				        width: auto !important;
				        height: auto !important
				    }
				    table[class="body"] center {
				        min-width: 0 !important
				    }
				    table[class="body"] .container {
				        width: 95% !important
				    }
				    table[class="body"] .row {
				        width: 100% !important;
				        display: block !important
				    }
        		}
        		@media all and (min-width: 48em) {
					blockquote {
						font-size: 34px;
						line-height: 40px;
						padding-top: 2px;
						padding-bottom: 3px;
					}
				}
        		table.two {
					width: 80px;
				}`)

	assert.Equal(t, len(css.CssRuleList), 4)

	assert.Equal(t, css.CssRuleList[0].Style.Selector.Text(), "table.one")
	assert.Equal(t, css.CssRuleList[0].Type, STYLE_RULE)
	assert.Equal(t, css.CssRuleList[0].Style.Styles[0].Value.Text(), "30px")
	assert.Equal(t, len(css.CssRuleList[0].Rules), 0)

	assert.Equal(t, css.CssRuleList[1].Style.Selector.Text(), "only screen and (max-width: 600px)")
	assert.Equal(t, css.CssRuleList[1].Type, MEDIA_RULE)
	assert.Equal(t, len(css.CssRuleList[1].Rules), 4)
	assert.Equal(t, css.CssRuleList[1].Rules[0].Style.Selector.Text(), "table[class=\"body\"] img")
	assert.Equal(t, css.CssRuleList[1].Rules[0].Style.Styles[1].Value.Text(), "auto")
	assert.Equal(t, css.CssRuleList[1].Rules[0].Style.Styles[1].Important, true)
	assert.Equal(t, css.CssRuleList[1].Rules[1].Style.Selector.Text(), "table[class=\"body\"] center")
	assert.Equal(t, css.CssRuleList[1].Rules[2].Style.Selector.Text(), "table[class=\"body\"] .container")
	assert.Equal(t, css.CssRuleList[1].Rules[3].Style.Selector.Text(), "table[class=\"body\"] .row")

	assert.Equal(t, css.CssRuleList[2].Style.Selector.Text(), "all and (min-width: 48em)")
	assert.Equal(t, css.CssRuleList[2].Type, MEDIA_RULE)
	assert.Equal(t, css.CssRuleList[2].Rules[0].Style.Selector.Text(), "blockquote")
	assert.Equal(t, css.CssRuleList[2].Rules[0].Style.Styles[0].Value.Text(), "34px")
	assert.Equal(t, css.CssRuleList[2].Rules[0].Style.Styles[1].Value.Text(), "40px")
	assert.Equal(t, css.CssRuleList[2].Rules[0].Style.Styles[2].Value.Text(), "2px")
	assert.Equal(t, css.CssRuleList[2].Rules[0].Style.Styles[3].Value.Text(), "3px")
	assert.Equal(t, len(css.CssRuleList[2].Rules), 1)

	assert.Equal(t, css.CssRuleList[3].Style.Selector.Text(), "table.two")
	assert.Equal(t, css.CssRuleList[3].Type, STYLE_RULE)
	assert.Equal(t, css.CssRuleList[3].Style.Styles[0].Value.Text(), "80px")
	assert.Equal(t, len(css.CssRuleList[3].Rules), 0)
}
