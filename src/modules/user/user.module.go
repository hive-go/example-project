package user

import (
	"github.com/hive-go/hive"
)

var UserModule = hive.CreateModule()

func init() {
	UserModule.AddController(UserController)
}
