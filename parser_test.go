package css

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithoutImpotant(t *testing.T) {
	css := Parse(`div .a { font-size: 150%;}`)
	assert.Equal(t, css.CssRuleList[0].Style.Styles["font-size"].Value, "150%")
	assert.Equal(t, css.CssRuleList[0].Style.Styles["font-size"].Property, "font-size")
	assert.Equal(t, css.CssRuleList[0].Style.Styles["font-size"].Important, 0)
	assert.Equal(t, css.CssRuleList[0].Style.SelectorText, "div .a")

}

func TestWithImpotant(t *testing.T) {
	css := Parse("div .a { font-size: 150% !important;}")
	assert.Equal(t, css.CssRuleList[0].Style.Styles["font-size"].Value, "150%")
	assert.Equal(t, css.CssRuleList[0].Style.Styles["font-size"].Property, "font-size")
	assert.Equal(t, css.CssRuleList[0].Style.Styles["font-size"].Important, 1)
	assert.Equal(t, css.CssRuleList[0].Style.SelectorText, "div .a")
}

func TestMultipleDeclarations(t *testing.T) {
	css := Parse(`div .a {
				font-size: 150%;
				width: 100%
				}`)
	assert.Equal(t, css.CssRuleList[0].Style.Styles["font-size"].Value, "150%")
	assert.Equal(t, css.CssRuleList[0].Style.Styles["font-size"].Property, "font-size")
	assert.Equal(t, css.CssRuleList[0].Style.Styles["font-size"].Important, 0)
	assert.Equal(t, css.CssRuleList[0].Style.Styles["width"].Value, "100%")
	assert.Equal(t, css.CssRuleList[0].Style.Styles["width"].Property, "width")
	assert.Equal(t, css.CssRuleList[0].Style.Styles["width"].Important, 0)
	assert.Equal(t, css.CssRuleList[0].Style.SelectorText, "div .a")
}

func TestMultipleSelectors(t *testing.T) {
	css := Parse(`div .a {
						font-size: 150%;
					}
					p .b {
						font-size: 250%;
					}`)

	assert.Equal(t, css.CssRuleList[0].Style.SelectorText, "div .a")
	assert.Equal(t, css.CssRuleList[1].Style.SelectorText, "p .b")

}

func TestValuePx(t *testing.T) {
	css := Parse("div .a { font-size: 45px;}")
	assert.Equal(t, css.CssRuleList[0].Style.Styles["font-size"].Value, "45px")
}

func TestValueEm(t *testing.T) {
	css := Parse("div .a { font-size: 45em;}")
	assert.Equal(t, css.CssRuleList[0].Style.Styles["font-size"].Value, "45em")
}

func TestValueHex(t *testing.T) {
	css := Parse("div .a { color: #123456;}")
	assert.Equal(t, css.CssRuleList[0].Style.Styles["color"].Value, "#123456")
}

func TestValueRGBFunction(t *testing.T) {
	css := Parse(".color{ color: rgb(1,2,3);}")

	assert.Equal(t, css.CssRuleList[0].Style.Styles["color"].Value, "rgb(1,2,3)")
	assert.Equal(t, css.CssRuleList[0].Style.SelectorText, ".color")
}

func TestValueString(t *testing.T) {
	css := Parse("div .center { text-align: center; }")

	assert.Equal(t, css.CssRuleList[0].Style.Styles["text-align"].Value, "center")
}

func TestId(t *testing.T) {
	css := Parse("#div { color: red;}")

	assert.Equal(t, css.CssRuleList[0].Style.Styles["color"].Value, "red")
	assert.Equal(t, css.CssRuleList[0].Style.SelectorText, "#div")
}

func TestClass(t *testing.T) {
	css := Parse(".div { color: green;}")

	assert.Equal(t, css.CssRuleList[0].Style.Styles["color"].Value, "green")
	assert.Equal(t, css.CssRuleList[0].Style.SelectorText, ".div")
}

func TestValueWhiteSpace(t *testing.T) {
	css := Parse(".div { padding: 10px 0 0 10px}")

	assert.Equal(t, "10px 0 0 10px", css.CssRuleList[0].Style.Styles["padding"].Value)
	assert.Equal(t, css.CssRuleList[0].Style.SelectorText, ".div")
}
