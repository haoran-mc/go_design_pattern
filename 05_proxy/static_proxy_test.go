package proxy

import (
	"testing"
)

func TestUserProxyLogin(t *testing.T) {
	proxy := NewUserProxy(&User{})

	if err := proxy.Login("test", "password"); err != nil {
		t.Error("failed to login")
	}
}
