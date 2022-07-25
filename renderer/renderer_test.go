package renderer

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func TestViewSymbol(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout
	ViewSymbol(os.Stdout)
	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	t.Run("should print > ", func(t *testing.T) {
		assert.Equal(t, "> ", string(bytes))
	})
}

func TestViewValue(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout
	ViewValue(os.Stdout, 2)
	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	t.Run("should print the value 2", func(t *testing.T) {
		assert.Equal(t, "2.0000\n", string(bytes))
	})
}

func TestViewError(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout
	ViewError(os.Stdout, "error")
	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	t.Run("should print given error", func(t *testing.T) {
		assert.Equal(t, "error\n", string(bytes))
	})
}

func setupStdout() func() {
	originalStdout := os.Stdout
	return func() {
		os.Stdout = originalStdout
	}
}
