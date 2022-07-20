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

func TestParser(t *testing.T) {
	t.Run("when parsing the string should not be empty", func(t *testing.T) {
		assert.Panics(t, func() {
			parser := NewParser()
			parser.Parse("")
		})
	})

	t.Run("should return ['add', '5'] for the input 'add 5'", func(t *testing.T) {
		parser := NewParser()
		parser.Parse("add 5")
		assert.Equal(t, []string{"add", "5"}, parser.tokens)
	})

	t.Run("should return ['subtract', '10'] for the input 'subtract 10'", func(t *testing.T) {
		parser := NewParser()
		parser.Parse("subtract 10")
		assert.Equal(t, []string{"subtract", "10"}, parser.tokens)
	})

	t.Run("should return ['multiply', '3'] for the input 'multiply 3'", func(t *testing.T) {
		parser := NewParser()
		parser.Parse("multiply 3")
		assert.Equal(t, []string{"multiply", "3"}, parser.tokens)
	})

	t.Run("should return ['divide', '2.5'] for the input 'divide 2.5'", func(t *testing.T) {
		parser := NewParser()
		parser.Parse("divide 2.5")
		assert.Equal(t, []string{"divide", "2.5"}, parser.tokens)
	})
}
