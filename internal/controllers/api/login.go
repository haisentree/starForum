package api

import (
	"github.com/kataras/iris/v12"
	"starForum/internal/global"
	"starForum/internal/global/form"
	msg "starForum/internal/global/message"
	"starForum/internal/services"
)

type LoginController struct {
	Ctx iris.Context
}

// 注册
func (c *LoginController) PostSignup() {
	resp := msg.NewCommonResponse(nil)

	req := form.SignupMsgReq{}
	c.Ctx.ReadJSON(&req)
	if err := global.Validate.Struct(req); err != nil {
		resp.Status = -1
		resp.Message = "表单校验失败"
		c.Ctx.JSON(resp.JsonCommonResponse())
		return
	}

	dealInfo := services.CreateUser(req)
	if dealInfo.Error != global.DealInfoSuccess {
		resp.Status = -1
		resp.Message = dealInfo.Message
		c.Ctx.JSON(resp.JsonCommonResponse())
		return
	}
	resp.Data = req
	c.Ctx.JSON(resp.JsonCommonResponse())
}

// 登录
