package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCalculator(t *testing.T) {
	t.Run("should return a new calculator, when initialised", func(t *testing.T) {
		assert.IsType(t, Calculator{}, NewCalculator())
	})

}

func TestCalculator_Add(t *testing.T) {
	calc := NewCalculator()
	t.Run("should assign 5 to the output, when added with 5", func(t *testing.T) {
		calc.Add(5)
		assert.Equal(t, 5.0, calc.output)
	})
}
