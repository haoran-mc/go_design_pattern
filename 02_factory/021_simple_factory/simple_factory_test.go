package factory

import (
	"reflect"
	"testing"
)

func TestNewConfigParser(t *testing.T) {
	tests := []struct {
		name string
		args string
		want ConfigParser
	}{
		{
			name: "json",
			args: "json",
			want: jsonConfigParser{},
		},
		{
			name: "yaml",
			args: "yaml",
			want: yamlConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfigParser(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
