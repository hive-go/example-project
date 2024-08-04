package user

import (
	"github.com/hive-go/hive"
)

type UserServiceT struct{}

var UserService = UserServiceT{}

func (u *UserServiceT) GetUser(id string) hive.Map {
	return hive.Map{
		"id":      id,
		"message": "User Retrieved",
		"user":    hive.Map{"name": "John Doe"},
	}
}

func (u *UserServiceT) CreateUser(c *hive.Ctx) string {
	return "User Created"
}

func (u *UserServiceT) UpdateUser(c *hive.Ctx) string {
	return "User Updated"
}

func (u *UserServiceT) DeleteUser(c *hive.Ctx) string {
	return "User Deleted"
}
