package api

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"starForum/internal/global"
	"starForum/internal/global/form"
	"starForum/internal/global/message"
	"starForum/internal/services"
	"strings"
)

type UserController struct {
	Ctx iris.Context
}

// =============================不需要登录====================================
// 注册信息以email为键,存储在cache中
func (c *UserController) PostSigupInfo() {
	resp := message.NewCommonResponse()
	// 1.校验表单
	req := form.SignupInfoMsgReq{}
	c.Ctx.ReadJSON(&req)
	if err := global.Validate.Struct(req); err != nil {
		resp.Status = message.ControllerError
		resp.Message = err.Error()
		c.Ctx.JSON(resp.JsonCommonResponse())
		return
	}
	// 2.执行逻辑函数
	respService := services.UserService.SigupInfoDeal(req)
	if respService.Status != message.SuccessStatus {
		resp.Status = respService.Status
		resp.Message = respService.Message
		c.Ctx.JSON(resp.JsonCommonResponse())
		return
	}

	c.Ctx.JSON(respService.JsonCommonResponse())
}

// 邮件验证码成功后，将存储在cache中的用户信息，转存到数据库中，并且响应token
func (c *UserController) PostSigupEmailVerify() {
	resp := message.NewCommonResponse()
	req := form.SigupEmailVerifyMsgReq{}
	c.Ctx.ReadJSON(&req)
	if err := global.Validate.Struct(req); err != nil {
		resp.Status = message.ControllerError
		resp.Message = err.Error()
		c.Ctx.JSON(resp.JsonCommonResponse())
		return
	}

	respService := services.UserService.SignupEmailVerify(req)
	if respService.Status != message.SuccessStatus {
		resp.Status = respService.Status
		resp.Message = respService.Message
		c.Ctx.JSON(resp.JsonCommonResponse())
		return
	}
	// 返回token
	c.Ctx.JSON(respService.JsonCommonResponse())
}

// 用户登录,返回token
func (c *UserController) PostLogin() {
	resp := message.NewCommonResponse()
	req := form.UserLoginMsgReq{}
	c.Ctx.ReadJSON(&req)
	if err := global.Validate.Struct(req); err != nil {
		resp.Status = message.ControllerError
		resp.Message = err.Error()
		c.Ctx.JSON(resp.JsonCommonResponse())
		return
	}
	respService := services.UserService.UserLogin(req)
	if respService.Status != message.SuccessStatus {
		c.Ctx.JSON(respService.JsonCommonResponse())
		return
	}

	c.Ctx.JSON(respService.JsonCommonResponse())
}

// 需要登录
func (c *UserController) GetCurrent() {
	resp := message.NewCommonResponse()
	// 从表头中获取token
	authorization := c.Ctx.GetHeader("Authorization")
	if len(authorization) == 0 {
		resp.Status = message.ControllerError
		resp.Message = "用户未登录"
		c.Ctx.JSON(resp.JsonCommonResponse())
		return
	}
	userToken, _ := strings.CutPrefix(authorization, "Bearer ")

	respService := services.UserService.GetCurrentUserByToken(userToken)
	if respService.Status != message.SuccessStatus {
		c.Ctx.JSON(respService.JsonCommonResponse())
		return
	}

	c.Ctx.JSON(respService.JsonCommonResponse())
}

func (c *UserController) GetLast() {

	err := global.EmailSender.SendSampleCode("1234", "1974733812@qq.com")
	if err != nil {
		fmt.Println("emial error", err)
	}

	resp := message.NewCommonResponse()
	c.Ctx.JSON(resp.JsonCommonResponse())
}

func (c *UserController) GetTest() {
	temp, _ := global.Cache.Get("1974733812@qq.com")
	fmt.Println(temp)

	resp := message.NewCommonResponse()
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
