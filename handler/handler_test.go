package handler

import (
	"github.com/stretchr/testify/assert"
	"maths/calculator"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	t.Run("should contain the operation, when registered", func(t *testing.T) {
		RegisterHandler("add", HandleAddition)
		assert.Equal(t, true, Registry["add"] != nil)
	})
}

func TestExecuteHandler(t *testing.T) {
	t.Run("should execute the operation, when executed", func(t *testing.T) {
		newCalculator := calculator.NewCalculator()
		ExecuteHandler(40, "add", newCalculator)
		assert.Equal(t, 40.0, newCalculator.GetValue())
	})
}
