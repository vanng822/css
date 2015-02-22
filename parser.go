package css

import (
	"fmt"
	"github.com/gorilla/css/scanner"
	"strings"
)

type State int

const (
	STATE_NONE State = iota
	STATE_SELECTOR
	STATE_DECLARE_BLOCK
	STATE_PROPERTY
	STATE_VALUE
)

type ParserContext struct {
	State           State
	NowSelectorText string
	NowProperty     string
	NowValue        string
	NowImportant    int
	NowRuleType     RuleType
	CurrentRule     *CSSRule
}

func Parse(csstext string) *CSSStyleSheet {
	context := &ParserContext{
		State:           STATE_NONE,
		NowSelectorText: "",
		NowProperty:     "",
		NowValue:        "",
		NowImportant:    0,
		NowRuleType:     STYLE_RULE,
	}

	css := &CSSStyleSheet{}
	css.CssRuleList = make([]*CSSRule, 0)
	s := scanner.New(csstext)

	for {
		token := s.Next()

		fmt.Printf("%s:'%s'\n", token.Type.String(), token.Value)

		if token.Type == scanner.TokenEOF || token.Type == scanner.TokenError {
			break
		}

		switch token.Type {
		case scanner.TokenAtKeyword:
			context.State = STATE_SELECTOR
			switch token.Value {
			case "@media":
				context.NowRuleType = MEDIA_RULE
			}

		case scanner.TokenString:
			if context.State == STATE_SELECTOR {
				context.NowSelectorText += token.Value
			}
		case scanner.TokenURI:
		case scanner.TokenUnicodeRange:
		case scanner.TokenCDO:
		case scanner.TokenCDC:
		case scanner.TokenComment:
		case scanner.TokenIdent:

			if context.State == STATE_NONE || context.State == STATE_SELECTOR {
				context.State = STATE_SELECTOR
				context.NowSelectorText += strings.TrimSpace(token.Value)
				break
			}

			if context.State == STATE_DECLARE_BLOCK {
				context.State = STATE_PROPERTY
				context.NowProperty = strings.TrimSpace(token.Value)
				break
			}

			if context.State == STATE_VALUE {
				if token.Value == "important" {
					context.NowImportant = 1
				} else {
					context.NowValue = token.Value + " "
				}
				break
			}

		case scanner.TokenDimension:
			fallthrough
		case scanner.TokenS:
			if context.State == STATE_SELECTOR {
				context.NowSelectorText += token.Value
			} else if context.State == STATE_VALUE {
				context.NowValue += token.Value
			}

		case scanner.TokenChar:
			if context.State == STATE_NONE {
				if token.Value != "{" {
					context.State = STATE_SELECTOR
					context.NowSelectorText += token.Value
					break
				}
			}
			if context.State == STATE_SELECTOR {
				if "{" == token.Value {
					context.State = STATE_DECLARE_BLOCK
					context.CurrentRule = NewRule(context.NowRuleType)
					context.CurrentRule.Style.SelectorText = strings.TrimSpace(context.NowSelectorText)
					break
				} else {
					context.NowSelectorText += token.Value
				}
				break
			}

			if context.State == STATE_PROPERTY {
				if token.Value == ":" {
					context.State = STATE_VALUE
				}
				break
			}

			if context.State == STATE_DECLARE_BLOCK {
				if token.Value == "}" {
					css.CssRuleList = append(css.CssRuleList, context.CurrentRule)
					context.NowSelectorText = ""
					context.NowProperty = ""
					context.NowValue = ""
					context.NowImportant = 0
					context.NowRuleType = STYLE_RULE
					context.State = STATE_NONE
				}
				break
			}

			if context.State == STATE_VALUE {
				if token.Value == ";" {
					decl := NewCSSStyleDeclaration(context.NowProperty, strings.TrimSpace(context.NowValue), context.NowImportant)
					context.CurrentRule.Style.Styles[context.NowProperty] = decl

					context.NowProperty = ""
					context.NowValue = ""
					context.NowImportant = 0
					context.State = STATE_DECLARE_BLOCK
				} else if token.Value == "}" { // last property in a block can have optional ;
					decl := NewCSSStyleDeclaration(context.NowProperty, strings.TrimSpace(context.NowValue), context.NowImportant)
					context.CurrentRule.Style.Styles[context.NowProperty] = decl
					css.CssRuleList = append(css.CssRuleList, context.CurrentRule)
					context.NowSelectorText = ""
					context.NowProperty = ""
					context.NowValue = ""
					context.NowImportant = 0
					context.NowRuleType = STYLE_RULE
					context.State = STATE_NONE
				} else if token.Value != "!" {
					context.NowValue += token.Value
				}
				break

			}
		case scanner.TokenPercentage:
			fallthrough
		case scanner.TokenHash:
			if context.State == STATE_NONE || context.State == STATE_SELECTOR {
				context.State = STATE_SELECTOR
				context.NowSelectorText += strings.TrimSpace(token.Value)
				break
			}

			fallthrough
		case scanner.TokenNumber:
			fallthrough
		case scanner.TokenFunction:
			fallthrough
		case scanner.TokenIncludes:
			fallthrough
		case scanner.TokenDashMatch:
			fallthrough
		case scanner.TokenPrefixMatch:
			fallthrough
		case scanner.TokenSuffixMatch:
			fallthrough
		case scanner.TokenSubstringMatch:
			if context.State == STATE_VALUE {
				context.NowValue += token.Value
			}
		default:
			fmt.Printf("Unhandled, %s:'%s'\n", token.Type.String(), token.Value)
		}
	}
	return css
}
