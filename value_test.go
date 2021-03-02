package css

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCSSValueString(t *testing.T) {
	val := NewCSSValueString("identifier")
	assert.Equal(t, val.Text(), `"identifier"`)

	val = NewCSSValueString(`with"char`)
	assert.Equal(t, val.Text(), `"with\\"char"`)
}
