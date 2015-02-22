package css

import ()

type CSSStyleDeclaration struct {
	Property  string
	Value     string
	Important int
}

func NewCSSStyleDeclaration(property, value string, important int) *CSSStyleDeclaration {
	return &CSSStyleDeclaration{
		Important: important,
		Value:     value,
		Property:  property,
	}
}
