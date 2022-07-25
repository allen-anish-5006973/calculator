package main

import (
	"maths/calculator"
	"maths/handler"
	"maths/parser"
	"maths/reader"
	"maths/renderer"
	"os"
)

func main() {
	calc := calculator.NewCalculator()
	read := os.Stdin
	writer := os.Stdout
	for true {
		renderer.ViewSymbol(writer)
		input := reader.Reader(read)
		tokenizer := parser.NewParser("", 0)
		tokenizer.Parse(input)
		operator, operand := tokenizer.GetTokens()
		handler.ExecuteHandler(operand, operator, calc)
	}
}
