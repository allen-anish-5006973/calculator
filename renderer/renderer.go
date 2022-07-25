package renderer

import (
	"fmt"
	"io"
	"strconv"
)

func ViewValue(writer io.Writer, value float64) {
	fmt.Fprintf(writer, strconv.FormatFloat(value, 'f', 4, 64))
	fmt.Println()
}

func ViewError(writer io.Writer, error string) {
	fmt.Fprintf(writer, error)
	fmt.Println()
}

func ViewSymbol(writer io.Writer) {
	fmt.Fprintf(writer, "> ")
}
