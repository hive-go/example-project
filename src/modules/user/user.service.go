package user

import (
  "github.com/gofiber/fiber/v2"
)

type UserServiceStruct struct{}

var UserService = UserServiceStruct{}

func (u *UserServiceStruct) CreateUser(data *CreateUserDto) (any, error) {
	return fiber.Map{
		"message": "User created successfully !!!",
	}, nil
}


func (u *UserServiceStruct) GetUsers() (any, error) {
  return fiber.Map{
    "message": "User retrieved successfully",
  }, nil
}

func (u *UserServiceStruct) GetUserById(id string) (any, error) {
  return fiber.Map{
    "message": "User retrieved successfully for id " + id,
  }, nil
}

func (u *UserServiceStruct) UpdateUser(id string, data *UpdateUserDto) (any, error) {
  return fiber.Map{
    "message": "User updated successfully for id " + id,
  }, nil
}

func (u *UserServiceStruct) DeleteUser(id string) (any, error) {
  return fiber.Map{
    "message": "User deleted successfully for id " + id,
  }, nil
}
