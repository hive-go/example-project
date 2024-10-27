package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hive-go/hive"
)

var UserController = hive.CreateController()

func init() {

	UserController.SetConfig(hive.ControllerConfig{
		Prefix: "/user",
		Tag:    "User",
	})

	UserController.ParseBody(CreateUserDto{}).Post("", func(c *fiber.Ctx) (any, error) {
		body := c.Locals("body").(*CreateUserDto)
		return UserService.CreateUser(body)
	})

	UserController.Get("", func(c *fiber.Ctx) (any, error) {
		return UserService.GetUsers()
	})

	UserController.Get("/:id", func(c *fiber.Ctx) (any, error) {
		id := c.Params("id")
		return UserService.GetUserById(id)
	})

	UserController.ParseBody(CreateUserDto{}).Patch("/:id", func(c *fiber.Ctx) (any, error) {
		id := c.Params("id")
		body := c.Locals("body").(*UpdateUserDto)
		return UserService.UpdateUser(id, body)
	})

	UserController.Delete("/:id", func(c *fiber.Ctx) (any, error) {
		id := c.Params("id")
		return UserService.DeleteUser(id)
	})
}
