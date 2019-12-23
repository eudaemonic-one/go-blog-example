package main

import (
	"errors"
	"strings"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userList = []user{
	user{Username: "user1", Password: "password1"},
	user{Username: "user2", Password: "password2"},
	user{Username: "user3", Password: "password3"},
}

func registerNewUser(username, password string) (*user, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty.")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("The username is unavailable.")
	}
	u := user{Username: username, Password: password}
	userList = append(userList, u)
	return &u, nil
}

func isUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}

func isUserValid(username, password string) bool {
	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}