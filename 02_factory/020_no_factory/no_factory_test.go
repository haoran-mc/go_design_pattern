package factory

import (
	"testing"
)

func TestGetFileExtension(t *testing.T) {
	jsonFileSuffix, _ := GetFileExtension("rule.json")
	if jsonFileSuffix != "json" {
		t.Error("expected: json, but: ", jsonFileSuffix)
	}

	_, err := GetFileExtension("rule.yaml")
	if err != nil {
		t.Error("expected: yaml, but: no ext")
	}
}

func TestLoad(t *testing.T) {
	Load("rule.json")
	Load("rule.yaml")
}
