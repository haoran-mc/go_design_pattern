package singleton_test

import (
	"testing"

	singleton "github.com/haoran-mc/go_design_pattern/01_singleton"
)

func TestGetInstance(t *testing.T) {
	singleton1 := singleton.GetSingletonInstance()
	singleton2 := singleton.GetSingletonInstance()
	if singleton1 != singleton2 {
		t.Error("expected: singleton1 == singleton2, but !=")
	}

	singleton3 := singleton.GetInstance()
	singleton4 := singleton.GetInstance()
	if singleton3 == singleton4 {
		t.Error("expected: singleton3 != singleton4, but ==")
	}
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetSingletonInstance() != singleton.GetSingletonInstance() {
				b.Error("test fail")
			}
		}
	})
}
