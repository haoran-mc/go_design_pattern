package main

import (
	"fmt"

	di "github.com/haoran-mc/go_design_pattern/02_factory/024_di"
)

type C struct {
	Num int
}

func NewC() *C {
	return &C{
		Num: 1,
	}
}

// B 依赖于 C
type B struct {
	C *C
}

func NewB(c *C) *B {
	return &B{C: c}
}

// A 依赖于 B
type A struct {
	B *B
}

func NewA(b *B) *A {
	return &A{
		B: b,
	}
}

func main() {
	container := di.New()
	if err := container.Provide(NewA); err != nil {
		panic(err)
	}
	if err := container.Provide(NewB); err != nil {
		panic(err)
	}
	if err := container.Provide(NewC); err != nil {
		panic(err)
	}

	err := container.Invoke(func(a *A) {
		fmt.Printf("%+v: %d", a, a.B.C.Num)
	})
	if err != nil {
		panic(err)
	}
}
