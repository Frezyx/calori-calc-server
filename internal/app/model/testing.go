package model

import "testing"

//TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "example@example.com",
		Password: "password",
	}
}

//TestUser2 ...
func TestUser2(t *testing.T) *User {
	return &User{
		Email:    "ex@example.com",
		Password: "123456",
	}
}
