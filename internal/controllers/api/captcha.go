package api

import (
	"github.com/kataras/iris/v12"
	"github.com/mojocn/base64Captcha"
	msg "starForum/internal/global/message"
)

type CaptchaController struct {
	Ctx iris.Context
}

var store = base64Captcha.DefaultMemStore
var digitDriver = base64Captcha.DriverDigit{
	Height:   50,
	Width:    200,
	Length:   4,   //验证码长度
	MaxSkew:  0.7, //倾斜
	DotCount: 1,   //背景的点数，越大，字体越模糊
}

func (c *CaptchaController) GetGenerate() {
	captcha := base64Captcha.NewCaptcha(&digitDriver, store)
	id, b64s, _, err := captcha.Generate()
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
	}
	resp := msg.NewCommonResponse(body)
	c.Ctx.JSON(resp.JsonCommonResponse())
}

type captchaReq struct {
	CaptchaId     string `json:"captchaId"`
	CaptchaAnswer string `json:"captchaAnswer"`
}

func (c *CaptchaController) PostVerify() {
	resp := msg.NewCommonResponse(nil)

	req := captchaReq{}
	c.Ctx.ReadJSON(&req)

	result := store.Verify(req.CaptchaId, req.CaptchaAnswer, true)
	resp.Data = iris.Map{
		"result": result,
	}
	c.Ctx.JSON(resp.JsonCommonResponse())
}
