package api

import (
	"github.com/kataras/iris/v12"
	"starForum/internal/global"
	msg "starForum/internal/global/message"
)

type CaptchaController struct {
	Ctx iris.Context
}

func (c *CaptchaController) GetGenerate() {
	//captcha := base64Captcha.NewCaptcha(&digitDriver, store)
	id, bs64, _, err := global.CaptchaGenerate.Generate()
	body := map[string]interface{}{"captchaId": id, "captchaBase64": bs64}
	if err != nil {
	}
	resp := msg.NewCommonResponse(body)
	c.Ctx.JSON(resp.JsonCommonResponse())
}

type captchaReq struct {
	CaptchaId     string `json:"captchaId"`
	CaptchaAnswer string `json:"captchaAnswer"`
}

// 这个接口在实际中不会使用
func (c *CaptchaController) PostVerify() {
	resp := msg.NewCommonResponse(nil)

	req := captchaReq{}
	c.Ctx.ReadJSON(&req)

	result := global.CaptchaStore.Verify(req.CaptchaId, req.CaptchaAnswer, true)
	resp.Data = iris.Map{
		"result": result,
	}
	c.Ctx.JSON(resp.JsonCommonResponse())
}
