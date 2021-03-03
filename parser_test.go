package css

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithoutImpotant(t *testing.T) {
	css := Parse(`div .a { font-size: 150%;}`)
	assert.Equal(t, css.CssRuleList[0].Style.Styles[0].Value.Text(), "150%")
	assert.Equal(t, css.CssRuleList[0].Style.Styles[0].Property, "font-size")
	assert.Equal(t, css.CssRuleList[0].Style.Styles[0].Important, false)
	assert.Equal(t, css.CssRuleList[0].Style.Selector.Text(), "div .a")

}

func TestWithImpotant(t *testing.T) {
	css := Parse("div .a { font-size: 150% !important;}")
	assert.Equal(t, css.CssRuleList[0].Style.Styles[0].Value.Text(), "150%")
	assert.Equal(t, css.CssRuleList[0].Style.Styles[0].Property, "font-size")
	assert.Equal(t, css.CssRuleList[0].Style.Styles[0].Important, true)
	assert.Equal(t, css.CssRuleList[0].Style.Selector.Text(), "div .a")
}

func TestMultipleDeclarations(t *testing.T) {
	css := Parse(`div .a {
				font-size: 150%;
				width: 100%
				}`)
	assert.Equal(t, css.CssRuleList[0].Style.Styles[0].Value.Text(), "150%")
	assert.Equal(t, css.CssRuleList[0].Style.Styles[0].Property, "font-size")
	assert.Equal(t, css.CssRuleList[0].Style.Styles[0].Important, false)
	assert.Equal(t, css.CssRuleList[0].Style.Styles[1].Value.Text(), "100%")
	assert.Equal(t, css.CssRuleList[0].Style.Styles[1].Property, "width")
	assert.Equal(t, css.CssRuleList[0].Style.Styles[1].Important, false)
	assert.Equal(t, css.CssRuleList[0].Style.Selector.Text(), "div .a")
}

func TestValuePx(t *testing.T) {
	css := Parse("div .a { font-size: 45px;}")
	assert.Equal(t, css.CssRuleList[0].Style.Styles[0].Value.Text(), "45px")
}

func TestValueEm(t *testing.T) {
	css := Parse("div .a { font-size: 45em;}")
	assert.Equal(t, css.CssRuleList[0].Style.Styles[0].Value.Text(), "45em")
}

func TestValueHex(t *testing.T) {
	css := Parse("div .a { color: #123456;}")
	assert.Equal(t, css.CssRuleList[0].Style.Styles[0].Value.Text(), "#123456")
}

func TestValueRGBFunction(t *testing.T) {
	css := Parse(".color{ color: rgb(1,2,3);}")

	assert.Equal(t, css.CssRuleList[0].Style.Styles[0].Value.Text(), "rgb(1,2,3)")
	assert.Equal(t, css.CssRuleList[0].Style.Selector.Text(), ".color")
}

func TestValueString(t *testing.T) {
	css := Parse("div .center { text-align: center; }")

	assert.Equal(t, css.CssRuleList[0].Style.Styles[0].Value.Text(), "center")
}

func TestValueWhiteSpace(t *testing.T) {
	css := Parse(".div { padding: 10px 0 0 10px}")

	assert.Equal(t, "10px 0 0 10px", css.CssRuleList[0].Style.Styles[0].Value.Text())
	assert.Equal(t, css.CssRuleList[0].Style.Selector.Text(), ".div")
}

func TestValueMixed(t *testing.T) {
	css := Parse(`td {
			padding: 0 12px 0 10px;
    		border-right: 1px solid white
		}`)

	assert.Equal(t, "0 12px 0 10px", css.CssRuleList[0].Style.Styles[0].Value.Text())
	assert.Equal(t, "1px solid white", css.CssRuleList[0].Style.Styles[1].Value.Text())
	assert.Equal(t, css.CssRuleList[0].Style.Selector.Text(), "td")
}

func TestQuoteValue(t *testing.T) {
	css := Parse(`blockquote {
    				font-family: "Source Sans Pro", Arial, sans-serif;
			    	font-size: 27px;
			    	line-height: 35px;}`)

	assert.Equal(t, "\"Source Sans Pro\", Arial, sans-serif", css.CssRuleList[0].Style.Styles[0].Value.Text())
	assert.Equal(t, "27px", css.CssRuleList[0].Style.Styles[1].Value.Text())
	assert.Equal(t, "35px", css.CssRuleList[0].Style.Styles[2].Value.Text())
	assert.Equal(t, css.CssRuleList[0].Style.Selector.Text(), "blockquote")
}

func TestDashClassname(t *testing.T) {
	css := Parse(`.content {
    				padding: 0px;
						}
						.content-wrap {
					  padding: 2px;
						}`)

	assert.Equal(t, ".content", css.CssRuleList[0].Style.Selector.Text())
	assert.Equal(t, ".content-wrap", css.CssRuleList[1].Style.Selector.Text())
	assert.Equal(t, "0px", css.CssRuleList[0].Style.Styles[0].Value.Text())
	assert.Equal(t, "2px", css.CssRuleList[1].Style.Styles[0].Value.Text())
}

func TestNotSupportedAtRule(t *testing.T) {
	rules := []string{
		`@namespace url(http://www.w3.org/1999/xhtml);`,
		`@document url(http://www.w3.org/),
               url-prefix(http://www.w3.org/Style/),
               domain(mozilla.org),
               regexp("https:.*")
			{

			  body { color: purple; background: yellow; }
			}`,
	}
	css := &CSSStyleSheet{}
	css.CssRuleList = make([]*CSSRule, 0)
	for _, rule := range rules {
		assert.Equal(t, css, Parse(rule))
	}
}
