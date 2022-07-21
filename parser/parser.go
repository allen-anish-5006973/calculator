package parser

import (
	"maths/constants"
	"maths/renderer"
	"strconv"
	"strings"
)

type Parser struct {
	operator constants.Operations
	operand  float64
}

func NewParser() *Parser {
	return &Parser{" ", 0}
}

func (parser *Parser) Parse(statement string) {
	tokens := strings.Split(statement, " ")
	parser.operator = constants.Operations(tokens[0])
	parser.operand = ParseOperand(tokens[1])
}

func ParseOperand(operand string) float64 {
	number, err := strconv.ParseFloat(operand, 64)
	if err != nil {
		renderer.ViewError("operand should be a number")
	}
	return number
}

func (parser Parser) GetTokens() (constants.Operations, float64) {
	return parser.operator, parser.operand
}
