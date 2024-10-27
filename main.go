package main

import (
	"example/src/modules/user"

	"github.com/gofiber/fiber/v2"
	"github.com/hive-go/hive"
)

func main() {
	app := hive.New(hive.Config{
		FiberConfig: fiber.Config{
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				code := fiber.StatusInternalServerError
				if e, ok := err.(*fiber.Error); ok {
					code = e.Code
				}
				return ctx.Status(code).JSON(fiber.Map{
					"message": err.Error(),
				})
			},
		},
		SwaggerConfig: hive.SwaggerConfig{
			Title:       "Aponta.ai API",
			Description: "Aponta.ai API",
			Version:     "1.0.0",
			Enabled:     true,
			Path:        "/swagger",
		},
	})

	app.AddModule(user.UserModule)

	app.Listen(":3000")
}
