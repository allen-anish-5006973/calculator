package renderer

import (
	"fmt"
	"io"
)

var writer io.Writer

func SetWriter(newWrite io.Writer) {
	writer = newWrite
}

func ViewValue(value float64) {
	fmt.Println(value)
}

func ViewError(error string) {
	fmt.Println(error)
}

func ViewSymbol() {
	fmt.Print("> ")
}
