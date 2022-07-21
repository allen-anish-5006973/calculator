package parser

import (
	"github.com/stretchr/testify/assert"
	"maths/constants"
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
		operator, operand := parser.GetTokens()
		assert.Equal(t, constants.Operations("add"), operator)
		assert.Equal(t, 5.0, operand)
	})

	t.Run("should return ['subtract', '10'] for the input 'subtract 10'", func(t *testing.T) {
		parser := NewParser()
		parser.Parse("subtract 10")
		operator, operand := parser.GetTokens()
		assert.Equal(t, constants.Operations("subtract"), operator)
		assert.Equal(t, 10.0, operand)
	})

	t.Run("should return ['multiply', '3'] for the input 'multiply 3'", func(t *testing.T) {
		parser := NewParser()
		parser.Parse("multiply 3")
		operator, operand := parser.GetTokens()
		assert.Equal(t, constants.Operations("multiply"), operator)
		assert.Equal(t, 3.0, operand)
	})

	t.Run("should return ['divide', '2.5'] for the input 'divide 2.5'", func(t *testing.T) {
		parser := NewParser()
		parser.Parse("divide 2.5")
		operator, operand := parser.GetTokens()
		assert.Equal(t, constants.Operations("divide"), operator)
		assert.Equal(t, 2.5, operand)
	})
}
