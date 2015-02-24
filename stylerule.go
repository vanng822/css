package css

import (
	"fmt"
	"sort"
	"strings"
)

type CSSStyleRule struct {
	SelectorText string
	Styles       map[string]*CSSStyleDeclaration
}

func (sr *CSSStyleRule) Text() string {
	decls := make([]string, 0, len(sr.Styles))
	keys := make([]string, 0, len(sr.Styles))
	for key, _ := range sr.Styles {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		decls = append(decls, sr.Styles[key].Text())
	}

	return fmt.Sprintf("%s {\n%s\n}", sr.SelectorText, strings.Join(decls, ";\n"))
}
