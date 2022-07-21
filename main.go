package main

import (
	"maths/calculator"
	"maths/handler"
	"maths/parser"
	"maths/reader"
	"os"
)

func main() {
	calc := calculator.NewCalculator()
	for true {
		input := reader.Reader(os.Stdin)
		tokenizer := parser.NewParser()
		tokenizer.Parse(input)
		operator, operand := tokenizer.GetTokens()
		handler.ExecuteHandler(operand, operator, calc)
	}
}
