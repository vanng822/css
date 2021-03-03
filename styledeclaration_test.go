package css

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeclWithImportan(t *testing.T) {
	decl := NewCSSStyleDeclaration("width", "100%", true)
	assert.Equal(t, decl.Text(), "width: 100% !important")
}

func TestDeclWithoutImportan(t *testing.T) {
	decl := NewCSSStyleDeclaration("width", "100%", false)
	assert.Equal(t, decl.Text(), "width: 100%")
}
