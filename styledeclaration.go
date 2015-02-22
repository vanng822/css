package css

import ()

type CSSStyleDeclaration struct {
	Property  string
	Value     string
	Important int
}

func NewCSSStyleDeclaration(property, value string, important int) *CSSStyleDeclaration {
	return &CSSStyleDeclaration{
		Property:  property,
		Value:     value,
		Important: important,
	}
}
