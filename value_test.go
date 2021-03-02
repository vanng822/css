package css

import (
	"testing"

	"github.com/gorilla/css/scanner"
	"github.com/stretchr/testify/assert"
)

func TestCSSValue(t *testing.T) {
	for _, c := range []struct{ inp, out string }{
		{`"identifier"`, `identifier`},
		{`identifier/path`, `identifier/path`},
		{`"with \' \" \\ chars"`, `with ' " \ chars`},
	} {
		t.Run(c.inp, func(t *testing.T) {
			val := NewCSSValue(c.inp)
			assert.Equal(t, val.Text(), c.inp)
			assert.Equal(t, val.ParsedText(), c.out)
		})
	}
}

func TestNewCSSValueString(t *testing.T) {
	for _, c := range []struct{ inp, out string }{
		{`identifier`, `"identifier"`},
		{`with special ' \ " char`, `"with special ' \\ \" char"`},
	} {
		t.Run(c.inp, func(t *testing.T) {
			val := NewCSSValueString(c.inp)
			assert.Equal(t, val.Text(), c.out)
			assert.Equal(t, val.ParsedText(), c.inp)
		})
	}
}

func TestSplitOnToken(t *testing.T) {
	val := NewCSSValue(`monospace, font-name, "font3",sans-serif`)
	split := val.SplitOnToken(&scanner.Token{Type: scanner.TokenChar, Value: ","})

	assert.Equal(t, len(split), 4)
	assert.Equal(t, split[0].ParsedText(), "monospace")
	assert.Equal(t, split[1].ParsedText(), "font-name")
	assert.Equal(t, split[2].ParsedText(), "font3")
	assert.Equal(t, split[3].ParsedText(), "sans-serif")
}
