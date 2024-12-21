package form

// ===================注册信息===============================
type SignupInfoMsgReq struct {
	CaptchaId     string `json:"captchaID" validate:"required"`
	CaptchaAnswer string `json:"captchaAnswer" validate:"required"`
	Email         string `json:"email"  validate:"required,email"`
	Username      string `json:"username"  validate:"required"`
	Password      string `json:"password"  validate:"required"`
	Nickname      string `json:"nickname"  validate:"required"`
	Avatar        string `json:"avatar"  validate:"required"`
}

// 缓存的注册信息
type SignupInfoMsgCache struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`

	EmailCode string `json:"emailCode"`
}

// 邮箱验证码验证信息
type SigupEmailVerifyMsgReq struct {
	Email     string `json:"email" validate:"required"`
	EmailCode string `json:"emailCode" validate:"required"`
}

// ============================登录信息==================================
type UserLoginMsgReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type UserLoginMsgResp struct {
	UserId   uint   `json:"userId"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`

	Token string `json:"token"  validate:"required"`
}

// 登录信息
type LoginMsgReq struct {
	Email    string `json:"email"  validate:"required"`
	Username string `json:"username"  validate:"required"`
	Password string `json:"password"  validate:"required"`
	Nickname string `json:"nickname"  validate:"required"`
	Avatar   string `json:"avatar"  validate:"required"`
}
type UserTokenMsgDealReq struct {
	Token      string
	UserId     uint
	ExpireTime int64
}

type UserTokenMsgDealResp struct {
	UserId uint
	Token  string
}

type CurrentUserMsgResp struct {
	UserID   uint   `json:"userID"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
