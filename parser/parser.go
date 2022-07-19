package parser

import (
	"strconv"
	"strings"
)

func Parse(statement string) (string, float64) {
	tokens := strings.Split(statement, " ")
	return tokens[0], strconv.ParseFloat(tokens[2], 64)
}
