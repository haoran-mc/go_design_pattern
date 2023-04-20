package proxy

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	want := `package proxy

type UserProxy struct {
	child *User
}

func NewUserProxy(child *User) *UserProxy {
	return &UserProxy{child: child}
}

func (p *UserProxy) Login(username, password string) (r0 error) {
	// before 这里可能会有一些统计的逻辑
	start := time.Now()

	r0 = p.child.Login(username, password)

	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return r0
}
`
	got, err := Generate("./static_proxy.go")
	if err != nil {
		t.Error("failed to generate")
	}
	if !reflect.DeepEqual(want, got) {
		t.Error("not equal")
	}
}
