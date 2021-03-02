package css

import (
	"testing"

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
