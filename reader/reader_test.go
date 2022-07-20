package reader

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	t.Run("when initialising, reader must return string", func(t *testing.T) {
		reader := strings.NewReader("")
		assert.IsType(t, "", Reader(reader))
	})

	t.Run("should return empty string when read by the reader", func(t *testing.T) {
		input := ""
		reader := strings.NewReader(input)
		assert.Equal(t, input, Reader(reader))
	})

	t.Run("should return string 'add 5' read by the reader", func(t *testing.T) {
		input := "add 5"
		reader := strings.NewReader(input)
		assert.Equal(t, input, Reader(reader))
	})

	t.Run("should return string 'subtract -10' read by the reader", func(t *testing.T) {
		input := "subtract -10"
		reader := strings.NewReader(input)
		assert.Equal(t, input, Reader(reader))
	})

	t.Run("should return string 'multiply 0.23' read by the reader", func(t *testing.T) {
		input := "multiply 0.23"
		reader := strings.NewReader(input)
		assert.Equal(t, input, Reader(reader))
	})

	t.Run("should return string 'divide -5.7' read by the reader", func(t *testing.T) {
		input := "divide -5.7"
		reader := strings.NewReader(input)
		assert.Equal(t, input, Reader(reader))
	})
}
