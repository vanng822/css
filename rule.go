package css

import ()

type RuleType int

const (
	STYLE_RULE RuleType = iota
	MEDIA_RULE
)

type CSSRule struct {
	Type  RuleType
	Style CSSStyleRule
}

func NewRule(ruleType RuleType) *CSSRule {
	r := &CSSRule{
		Type: ruleType,
	}
	r.Style.Styles = make(map[string]*CSSStyleDeclaration)

	return r
}
