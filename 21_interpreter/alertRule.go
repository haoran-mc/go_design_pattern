package interpreter

// 告警规则
type alertRule struct {
	exp expression
}

// 判断告警是否触发
func (r alertRule) interpret(stats map[string]float64) bool {
	return r.exp.interpret(stats)
}

func newAlertRule(rule string) (*alertRule, error) {
	exp, err := newAndExpression(rule)
	return &alertRule{exp: exp}, err
}
