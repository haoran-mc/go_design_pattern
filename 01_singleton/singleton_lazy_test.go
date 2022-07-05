package singleton_test

import (
	singleton "github.com/haoran-mc/go_design_pattern/01_singleton"
	"testing"
)

func TestGetLazyInstance(t *testing.T) {
	singleton1 := singleton.GetLazySingletonInstance()
	singleton2 := singleton.GetLazySingletonInstance()
	if singleton1 != singleton2 {
		t.Error("expected: singleton1 == singleton2, but !=")
	}

	singleton3 := singleton.GetLazyInstance()
	singleton4 := singleton.GetLazyInstance()
	if singleton3 == singleton4 {
		t.Error("expected: singleton3 != singleton4, but ==")
	}
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetLazySingletonInstance() != singleton.GetLazySingletonInstance() {
				b.Error("test fail")
			}
		}
	})
}
