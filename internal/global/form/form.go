package form

type SignupMsgReq struct {
	Email    string `json:"email"  validate:"required"`
	Username string `json:"username"  validate:"required"`
	Password string `json:"password"  validate:"required"`
	Nickname string `json:"nickname"  validate:"required"`
	Avatar   string `json:"avatar"  validate:"required"`
}

type UserTokenMsgReq struct {
	UserId uint64 `json:"userId"  validate:"required"`
}

type UserTokenMsgDeal struct {
	Token  string
	UserId uint64
}