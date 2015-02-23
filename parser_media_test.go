package css

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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
	assert.Equal(t, css.CssRuleList[0].Style.SelectorText, "only screen and (max-width: 600px)")
	assert.Equal(t, css.CssRuleList[0].Type, MEDIA_RULE)
	assert.Equal(t, len(css.CssRuleList[0].Rules), 4)
	assert.Equal(t, css.CssRuleList[0].Rules[0].Style.SelectorText, "table[class=\"body\"] img")
	assert.Equal(t, css.CssRuleList[0].Rules[0].Style.Styles["height"].Value, "auto")
	assert.Equal(t, css.CssRuleList[0].Rules[0].Style.Styles["height"].Important, 1)
	assert.Equal(t, css.CssRuleList[0].Rules[1].Style.SelectorText, "table[class=\"body\"] center")
	assert.Equal(t, css.CssRuleList[0].Rules[2].Style.SelectorText, "table[class=\"body\"] .container")
	assert.Equal(t, css.CssRuleList[0].Rules[3].Style.SelectorText, "table[class=\"body\"] .row")

}

