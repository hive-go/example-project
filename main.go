package main

import (
	"example/src/modules/user"

	"github.com/hive-go/hive"
)

func main() {
	app := hive.New()

	app.AddModule(user.UserModule)

	app.Listen(":3000")
}
