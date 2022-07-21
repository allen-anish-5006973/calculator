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
	writer := os.Stdin
	for true {
		renderer.ViewSymbol()
		input := reader.Reader(writer)
		tokenizer := parser.NewParser()
		tokenizer.Parse(input)
		operator, operand := tokenizer.GetTokens()
		handler.ExecuteHandler(operand, operator, calc)
	}
}
