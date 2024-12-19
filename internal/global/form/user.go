package form

// 注册信息
type SignupInfoMsgReq struct {
	CaptchaId     string `json:"captchaID" validate:"required"`
	CaptchaAnswer string `json:"captchaAnswer" validate:"required"`
	Email         string `json:"email"  validate:"required,email"`
	Username      string `json:"username"  validate:"required"`
	Password      string `json:"password"  validate:"required"`
	Nickname      string `json:"nickname"  validate:"required"`
	Avatar        string `json:""  validate:"required"`
}

// 登录信息
type LoginMsgReq struct {
	Email    string `json:"email"  validate:"required"`
	Username string `json:"username"  validate:"required"`
	Password string `json:"password"  validate:"required"`
	Nickname string `json:"nickname"  validate:"required"`
	Avatar   string `json:"avatar"  validate:"required"`
}

type UserTokenMsgReq struct {
	UserId uint `json:"userId"  validate:"required"`
}

type UserTokenMsgDeal struct {
	Token      string
	UserId     uint
	ExpireTime int64
}
