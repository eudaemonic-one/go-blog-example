package main

import "testing"

func TestValidUserRegistration(t *testing.T) {
	saveUserList()
	u, err := registerNewUser("newuser", "newpassword")
	if err != nil || u.Username == "" {
		t.Fail()
	}
	restoreUserList()
}

func TestInvalidUserRegistration(t *testing.T) {
	saveUserList()
	u, err := registerNewUser("user1", "pass1")
	if err == nil || u != nil {
		t.Fail()
	}
	u, err = registerNewUser("newuser", "")
	if err == nil || u != nil {
		t.Fail()
	}
	restoreUserList()
}

func TestUsernameAvailablility(t *testing.T) {
	saveUserList()
	if !isUsernameAvailable("newuser") {
		t.Fail()
	}
	if isUsernameAvailable("user1") {
		t.Fail()
	}
	registerNewUser("newuser", "newpass")
	if isUsernameAvailable("newuser") {
		t.Fail()
	}
	restoreUserList()
}