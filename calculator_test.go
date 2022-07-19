package maths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCalculator(t *testing.T) {
	t.Run("should return a new calculator, when initialised", func(t *testing.T) {
		assert.IsType(t, Calculator{}, NewCalculator())
	})

}
