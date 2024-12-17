package form

type SignupMsgReq struct {
	Email    string `json:"email"  validate:"required"`
	Username string `json:"username"  validate:"required"`
	Password string `json:"password"  validate:"required"`
	Nickname string `json:"nickname"  validate:"required"`
	Avatar   string `json:"avatar"  validate:"required"`
}
