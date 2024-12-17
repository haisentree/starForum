package api

import (
	"github.com/kataras/iris/v12"
	msg "starForum/internal/global/message"
)

type UserController struct {
	Ctx iris.Context
}

func (c *UserController) GetCurrent() {
	resp := msg.NewCommonResponse("current")

	c.Ctx.JSON(resp.JsonCommonResponse())
}

func (c *UserController) GetLast() {
	resp := msg.NewCommonResponse("last")
	c.Ctx.JSON(resp.JsonCommonResponse())
}
