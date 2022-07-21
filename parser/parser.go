package parser

import (
	"maths/constants"
	"strconv"
	"strings"
)

type Parser struct {
	operator constants.Operations
	operand  float64
}

func NewParser() Parser {
	return Parser{}
}

func (parser *Parser) Parse(statement string) {
	if statement == "" {
		panic("Cannot parse an empty string")
	}
	tokens := strings.Split(statement, " ")
	if len(tokens) > 2 {
		panic("Only operator and operand is expected")
	}

	parser.operator = constants.Operations(tokens[0])
	if len(tokens) < 2 {
		parser.operand = 0
		return
	}

	parser.operand = ParseOperand(tokens[1])
}

func ParseOperand(operand string) float64 {
	number, err := strconv.ParseFloat(operand, 64)
	if err != nil {
		panic("operand should be a number")
	}
	return number
}

func (parser Parser) GetTokens() (constants.Operations, float64) {
	return parser.operator, parser.operand
}
