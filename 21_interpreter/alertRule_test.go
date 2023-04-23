package interpreter

import (
	"reflect"
	"testing"
)

func TestAlertRuleInterpret(t *testing.T) {
	stats := map[string]float64{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	tests := []struct {
		name  string
		stats map[string]float64
		rule  string
		want  bool
	}{
		{
			name:  "case1",
			stats: stats,
			rule:  "a > 1 && b > 10 && c < 5",
			want:  false,
		},
		{
			name:  "case2",
			stats: stats,
			rule:  "a < 2 && b > 10 && c < 5",
			want:  false,
		},
		{
			name:  "case3",
			stats: stats,
			rule:  "a < 5 && b > 1 && c < 10",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := newAlertRule(tt.rule)
			if err != nil {
				t.Error("failed to create alert rule")
			}
			if !reflect.DeepEqual(tt.want, r.interpret(tt.stats)) {
				t.Error("not equal")
			}
		})
	}
}
