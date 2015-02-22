package css

import ()

type CSSStyleRule struct {
	SelectorText string
	Styles       map[string]*CSSStyleDeclaration
}
