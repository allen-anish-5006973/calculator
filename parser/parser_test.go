package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewParser(t *testing.T) {
	t.Run("when initializing parser, should return a parser", func(t *testing.T) {
		assert.IsType(t, Parser{}, NewParser())
	})
}
