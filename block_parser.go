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

func ParseBlock(csstext string) map[string]*CSSStyleDeclaration {
	s := scanner.New(csstext)
	return parseBlock(s)
}

func parseBlock(s *scanner.Scanner) map[string]*CSSStyleDeclaration {
	/* block       : '{' S* [ any | block | ATKEYWORD S* | ';' S* ]* '}' S*;
	property    : IDENT;
	value       : [ any | block | ATKEYWORD S* ]+;
	any         : [ IDENT | NUMBER | PERCENTAGE | DIMENSION | STRING
	              | DELIM | URI | HASH | UNICODE-RANGE | INCLUDES
	              | DASHMATCH | ':' | FUNCTION S* [any|unused]* ')'
	              | '(' S* [any|unused]* ')' | '[' S* [any|unused]* ']'
	              ] S*;
	*/
	decls := make(map[string]*CSSStyleDeclaration)

	context := &ParserContext{
		State:           STATE_DECLARE_BLOCK,
		NowSelectorText: "",
		NowProperty:     "",
		NowValue:        "",
		NowImportant:    0,
	}

	for {
		token := s.Next()

		fmt.Printf("BLOCK(%d): %s:'%s'\n", context.State, token.Type.String(), token.Value)

		if token.Type == scanner.TokenEOF || token.Type == scanner.TokenError {
			break
		}

		switch token.Type {

		case scanner.TokenS:
			if context.State == STATE_VALUE {
				context.NowValue += token.Value
			}
		case scanner.TokenIdent:
			if context.State == STATE_DECLARE_BLOCK || context.State == STATE_NONE {
				context.State = STATE_PROPERTY
				context.NowProperty = strings.TrimSpace(token.Value)
				break
			}
			if token.Value == "important" {
				context.NowImportant = 1
			} else {
				context.NowValue += token.Value
			}
		case scanner.TokenChar:
			if context.State == STATE_PROPERTY {
				if token.Value == ":" {
					context.State = STATE_VALUE
				}
				// CHAR and STATE_PROPERTY but not : then weird
				// break to ignore it
				break
			}
			// should be no state or value
			if token.Value == ";" {
				decl := NewCSSStyleDeclaration(context.NowProperty, strings.TrimSpace(context.NowValue), context.NowImportant)
				decls[context.NowProperty] = decl
				context.NowProperty = ""
				context.NowValue = ""
				context.NowImportant = 0
				context.State = STATE_NONE
			} else if token.Value == "}" { // last property in a block can have optional ;
				if context.State == STATE_VALUE {
					// only valid if state is still VALUE, could be ;}
					decl := NewCSSStyleDeclaration(context.NowProperty, strings.TrimSpace(context.NowValue), context.NowImportant)
					decls[context.NowProperty] = decl
				}
				// we are done
				return decls
			} else if token.Value != "!" {
				context.NowValue += token.Value
			}
			break

		// any
		case scanner.TokenNumber:
			fallthrough
		case scanner.TokenPercentage:
			fallthrough
		case scanner.TokenDimension:
			fallthrough
		case scanner.TokenString:
			fallthrough
		case scanner.TokenURI:
			fallthrough
		case scanner.TokenHash:
			fallthrough
		case scanner.TokenUnicodeRange:
			fallthrough
		case scanner.TokenIncludes:
			fallthrough
		case scanner.TokenDashMatch:
			fallthrough
		case scanner.TokenFunction:
			fallthrough
		case scanner.TokenSubstringMatch:
			context.NowValue += token.Value
		}
	}

	return decls
}
