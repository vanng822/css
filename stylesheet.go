package css

import ()

type CSSStyleSheet struct {
	Type        string
	Media       string
	CssRuleList []*CSSRule
}
