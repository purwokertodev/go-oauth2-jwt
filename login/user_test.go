package login

import (
	"testing"
)

func TestSuccessLogin(t *testing.T) {
	u := &UserLogin{Username: "Wuriyanto", Password: "123456"}
	if !u.IsValidUser() {
		t.Error("Login is invalid")
	}
}

func TestFailLogin(t *testing.T) {
	u := &UserLogin{Username: "Wuriyanto", Password: "12345"}
	if u.IsValidUser() {
		t.Error("Login is invalid")
	}
}
