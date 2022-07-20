package parser

import "strings"

type Parser struct {
	tokens []string
}

func NewParser() Parser {
	return Parser{}
}

func (parser *Parser) Parse(statement string) {
	if statement == "" {
		panic("Cannot parse an empty string")
	}
	parser.tokens = strings.Split(statement, " ")
}
