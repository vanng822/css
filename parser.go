package css

import (
	"fmt"
	"github.com/gorilla/css/scanner"
	"strings"
)

/*
	stylesheet  : [ CDO | CDC | S | statement ]*;
	statement   : ruleset | at-rule;
	at-rule     : ATKEYWORD S* any* [ block | ';' S* ];
	block       : '{' S* [ any | block | ATKEYWORD S* | ';' S* ]* '}' S*;
	ruleset     : selector? '{' S* declaration? [ ';' S* declaration? ]* '}' S*;
	selector    : any+;
	declaration : property S* ':' S* value;
	property    : IDENT;
	value       : [ any | block | ATKEYWORD S* ]+;
	any         : [ IDENT | NUMBER | PERCENTAGE | DIMENSION | STRING
	              | DELIM | URI | HASH | UNICODE-RANGE | INCLUDES
	              | DASHMATCH | ':' | FUNCTION S* [any|unused]* ')'
	              | '(' S* [any|unused]* ')' | '[' S* [any|unused]* ']'
	              ] S*;
	unused      : block | ATKEYWORD S* | ';' S* | CDO S* | CDC S*;
*/

type State int

const (
	STATE_NONE State = iota
	STATE_SELECTOR
	STATE_PROPERTY
	STATE_VALUE
)

type parserContext struct {
	State            State
	NowSelectorText  string
	NowRuleType      RuleType
	CurrentRule      *CSSRule
	CurrentMediaRule *CSSRule
}

// Parse takes a string of valid css rules, stylesheet,
// and parses it. Be aware this function has poor error handling
// so you should have valid syntax in your css
func Parse(csstext string) *CSSStyleSheet {
	context := &parserContext{
		State:            STATE_NONE,
		NowSelectorText:  "",
		NowRuleType:      STYLE_RULE,
		CurrentMediaRule: nil,
	}

	css := &CSSStyleSheet{}
	css.CssRuleList = make([]*CSSRule, 0)
	s := scanner.New(csstext)

	for {
		token := s.Next()

		fmt.Printf("Parse(%d): %s:'%s'\n", context.State, token.Type.String(), token.Value)

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
				break
			}
		case scanner.TokenIdent:

			if context.State == STATE_NONE || context.State == STATE_SELECTOR {
				context.State = STATE_SELECTOR
				context.NowSelectorText += strings.TrimSpace(token.Value)
				break
			}
			
		case scanner.TokenDimension:
			fallthrough
		case scanner.TokenS:
			if context.State == STATE_SELECTOR {
				context.NowSelectorText += token.Value
				break
			}
			
		case scanner.TokenChar:
			if context.State == STATE_NONE {
				if token.Value == "}" && context.CurrentMediaRule != nil {
					// close media rule
					css.CssRuleList = append(css.CssRuleList, context.CurrentMediaRule)
					context.CurrentMediaRule = nil
					break
				}
				if token.Value != "{" {
					context.State = STATE_SELECTOR
					context.NowSelectorText += token.Value
					break
				}
			}
			if context.State == STATE_SELECTOR {
				if "{" == token.Value {
					if context.NowRuleType == MEDIA_RULE {
						context.CurrentMediaRule = NewRule(context.NowRuleType)
						context.CurrentMediaRule.Style.SelectorText = strings.TrimSpace(context.NowSelectorText)
						// reset
						context.NowSelectorText = ""
						context.NowRuleType = STYLE_RULE
						context.State = STATE_NONE
						break
					} else {
						context.CurrentRule = NewRule(context.NowRuleType)
						context.CurrentRule.Style.SelectorText = strings.TrimSpace(context.NowSelectorText)
						context.CurrentRule.Style.Styles = parseBlock(s)
						if context.CurrentMediaRule != nil {
							context.CurrentMediaRule.Rules = append(context.CurrentMediaRule.Rules, context.CurrentRule)
						} else {
							css.CssRuleList = append(css.CssRuleList, context.CurrentRule)
						}
						context.CurrentRule = nil
						context.NowSelectorText = ""
						context.NowRuleType = STYLE_RULE
						context.State = STATE_NONE
						break
					}
				} else {
					context.NowSelectorText += token.Value
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
			if context.State == STATE_SELECTOR {
				context.NowSelectorText += token.Value
				break
			}

		// Unhandle token
		case scanner.TokenURI:
			fallthrough
		case scanner.TokenUnicodeRange:
			fallthrough
		case scanner.TokenCDO:
			fallthrough
		case scanner.TokenCDC:
			fallthrough
		case scanner.TokenComment:
			fallthrough
		default:
			fmt.Printf("Unhandled, %s:'%s'\n", token.Type.String(), token.Value)
		}
	}
	return css
}
