package api

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"starForum/internal/global"
	"starForum/internal/global/form"
	msg "starForum/internal/global/message"
	"starForum/pkg/password"
)

type UserController struct {
	Ctx iris.Context
}

// 注册信息以email为键,存储在cache中
func (c *UserController) PostSigupInfo() {
	resp := msg.NewCommonResponse(nil)
	// 1.校验表单
	req := form.SignupInfoMsgReq{}
	c.Ctx.ReadJSON(&req)
	if err := global.Validate.Struct(req); err != nil {
		resp.Status = -1
		resp.Message = err.Error()
		c.Ctx.JSON(resp.JsonCommonResponse())
		return
	}
	// 2.执行逻辑函数

	resp.Data = req
	c.Ctx.JSON(resp.JsonCommonResponse())
}

// 不需要登录
//func (c *UserController) PostLogin() {}
//
//func (c *UserController) PostSigup() {
//
//}

// 需要登录
func (c *UserController) GetCurrent() {
	resp := msg.NewCommonResponse("current")
	pwd := "123456"
	enpwd := password.EncodePassword(pwd)
	res := password.ValidatePassword(enpwd, "1236456")
	fmt.Println("pwd:", pwd)
	fmt.Println("enpwd:", enpwd)
	fmt.Println("res:", res)

	c.Ctx.JSON(resp.JsonCommonResponse())
}

func (c *UserController) GetLast() {
	resp := msg.NewCommonResponse("last")
	c.Ctx.JSON(resp.JsonCommonResponse())
}

//func (c *UserController) PostLogout() {
//
//}

// 后台管理接口
//func (c *UserController) GetDealInfoByID() {
//
//}
//
//func (c *UserController) GetSimpleInfoByID() {
//
//}
