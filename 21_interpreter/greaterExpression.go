package interpreter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// >
type greaterExpression struct {
	key   string
	value float64
}

func (g greaterExpression) interpret(stats map[string]float64) bool {
	v, ok := stats[g.key]
	if !ok {
		return false
	}
	return v > g.value
}

func newGreaterExpression(exp string) (*greaterExpression, error) {
	data := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(exp), -1)
	if len(data) != 3 || data[1] != ">" {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	val, err := strconv.ParseFloat(data[2], 10)
	if err != nil {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	return &greaterExpression{
		key:   data[0],
		value: val,
	}, nil
}
