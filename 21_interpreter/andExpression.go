package interpreter

import (
	"fmt"
	"strings"
)

// &&
type andExpression struct {
	exps []expression
}

func (e andExpression) interpret(stats map[string]float64) bool {
	for _, expression := range e.exps {
		if !expression.interpret(stats) {
			return false
		}
	}
	return true
}

func newAndExpression(ruleStr string) (*andExpression, error) {
	rules := strings.Split(ruleStr, "&&") // []string{"a > 1", ...}
	exps := make([]expression, len(rules))

	for i, e := range rules {
		var exp expression
		var err error

		switch {
		case strings.Contains(e, ">"):
			exp, err = newGreaterExpression(e)
		case strings.Contains(e, "<"):
			exp, err = newLessExpression(e)
		default:
			err = fmt.Errorf("exp is invalid: %s", ruleStr)
		}

		if err != nil {
			return nil, err
		}

		exps[i] = exp
	}

	return &andExpression{exps: exps}, nil
}
