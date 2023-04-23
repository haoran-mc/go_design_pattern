package interpreter

// 表达式
type expression interface {
	interpret(stats map[string]float64) bool
}
