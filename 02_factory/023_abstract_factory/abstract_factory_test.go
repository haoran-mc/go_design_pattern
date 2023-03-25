package factory

import (
	"reflect"
	"testing"
)

func Test_jsonConfigParserFactory_CreateCarParser(t *testing.T) {
	tests := []struct {
		name string
		want CarParser
	}{
		{
			name: "json",
			want: jsonCarParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jsonParserFactory{}
			if got := j.CreateCarParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCarParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonConfigParserFactory_CreateComputerParser(t *testing.T) {
	tests := []struct {
		name string
		want ComputerParser
	}{
		{
			name: "json",
			want: jsonComputerParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jsonParserFactory{}
			if got := j.CreateComputerParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateComputerParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
