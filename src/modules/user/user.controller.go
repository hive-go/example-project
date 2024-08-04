package user

import (
	"github.com/hive-go/hive"
)

var UserController = hive.CreateController()

func init() {
	UserController.Get("/user", func(c *hive.Ctx) any {
		return UserService.GetUser("123")
	})

	UserController.Post("/user", func(c *hive.Ctx) any {
		return UserService.CreateUser(c)
	})

	UserController.Put("/user", func(c *hive.Ctx) any {
		return UserService.UpdateUser(c)
	})

	UserController.Delete("/user", func(c *hive.Ctx) any {
		return UserService.DeleteUser(c)
	})
}
