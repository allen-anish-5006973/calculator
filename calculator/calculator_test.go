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

	t.Run("should assign 12.5 to the output, when added with 7.5", func(t *testing.T) {
		calc.Add(7.5)
		assert.Equal(t, 12.5, calc.output)
	})
}

func TestCalculator_Subtract(t *testing.T) {
	calc := NewCalculator()
	t.Run("should assign -3 to the output, when subtracted with 3", func(t *testing.T) {
		calc.Subtract(3)
		assert.Equal(t, -3.0, calc.output)
	})

	t.Run("should assign -7.42 to the output, when subtracted with 4.42", func(t *testing.T) {
		calc.Subtract(4.42)
		assert.Equal(t, -7.42, calc.output)
	})
}

func TestCalculator_Multiply(t *testing.T) {
	calc := NewCalculator()
	calc.Add(1)
	t.Run("should assign 5 to the output, when multiplied with 5", func(t *testing.T) {
		calc.Multiply(5)
		assert.Equal(t, 5.0, calc.output)
	})

	t.Run("should assign -27.5 to the output, when multiplied with 5.5", func(t *testing.T) {
		calc.Multiply(-5.5)
		assert.Equal(t, -27.5, calc.output)
	})
}
