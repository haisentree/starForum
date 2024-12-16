package api

import (
	"github.com/kataras/iris/v12"
)

type UserController struct {
	Ctx iris.Context
}

func (c *UserController) GetCurrent() {
	c.Ctx.JSON(iris.Map{"status": iris.StatusOK,
		"message": "hello",
		"data":    "go",
	})
}
