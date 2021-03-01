package css

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBlock(t *testing.T) {
	css := ParseBlock(`
    				font-family: "Source Sans Pro", Arial, sans-serif;
			    	font-size: 27px;
			    	line-height: 35px;`)

	assert.Equal(t, len(css), 3)
	assert.Equal(t, "35px", css[2].Value.Text())
}

func TestParseBlockOneLine(t *testing.T) {
	css := ParseBlock("font-family: \"Source Sans Pro\", Arial, sans-serif; font-size: 27px;")

	assert.Equal(t, len(css), 2)
	assert.Equal(t, "27px", css[1].Value.Text())
	assert.Equal(t, "\"Source Sans Pro\", Arial, sans-serif", css[0].Value.Text())
}

func TestParseBlockBlankEnd(t *testing.T) {
	css := ParseBlock("font-size: 27px; width: 10px")

	assert.Equal(t, len(css), 2)
	assert.Equal(t, "27px", css[0].Value.Text())
	assert.Equal(t, "10px", css[1].Value.Text())
}

func TestParseBlockInportant(t *testing.T) {
	css := ParseBlock("font-size: 27px; width: 10px !important")

	assert.Equal(t, len(css), 2)
	assert.Equal(t, "27px", css[0].Value.Text())
	assert.Equal(t, "10px", css[1].Value.Text())
	assert.Equal(t, true, css[1].Important)
}

func TestParseBlockWithBraces(t *testing.T) {
	css := ParseBlock("{ font-size: 27px; width: 10px }")

	assert.Equal(t, len(css), 2)
	assert.Equal(t, "27px", css[0].Value.Text())
	assert.Equal(t, "10px", css[1].Value.Text())
}
