package singleton_test

import (
	singleton "github.com/haoran-mc/go_design_pattern/01_singleton"
	"testing"
)

func TestGetLazyInstance(t *testing.T) {
	singleton1 := singleton.GetLazyInstance()
	singleton2 := singleton.GetLazyInstance()
	if singleton1 != singleton2 {
		t.Error("expected: singleton1 == singleton2, but !=")
	}

	singleton3 := &singleton.Singleton{}
	singleton4 := &singleton.Singleton{}
	if singleton3 == singleton4 {
		t.Error("expected: singleton3 != singleton4, but ==")
	}
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetLazyInstance() != singleton.GetLazyInstance() {
				b.Error("test fail")
			}
		}
	})
}
