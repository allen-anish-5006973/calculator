package parser

import (
	"maths/constants"
	"maths/renderer"
	"os"
	"strconv"
	"strings"
)

type Parser struct {
	Operator constants.Operations
	Operand  float64
}

func NewParser(operator constants.Operations, operand float64) *Parser {
	return &Parser{operator, operand}
}

func (parser *Parser) Parse(statement string) {
	tokens := strings.Split(statement, " ")
	if len(tokens) >= 1 {
		parser.Operator = constants.Operations(tokens[0])
	}
	if len(tokens) >= 2 {
		parser.Operand = ParseOperand(tokens[1])
	}
}

func ParseOperand(operand string) float64 {
	number, err := strconv.ParseFloat(operand, 64)
	if err != nil {
		renderer.ViewError(os.Stdout, "Operand should be a number")
	}
	return number
}

func (parser Parser) GetTokens() (constants.Operations, float64) {
	return parser.Operator, parser.Operand
}
