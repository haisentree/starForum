package models

import (
	"gorm.io/gorm"
	"starForum/internal/global"
	"starForum/internal/global/form"
	"starForum/internal/global/message"
)

// ========================================UserToken========================================
type UserToken struct {
	gorm.Model
	Token       string `gorm:"unique;not null" json:"token"`
	UserId      uint   `gorm:"not null;index:idx_user_token_user_id;" json:"userId"`
	ExpiredTime int64  `gorm:"not null" json:"expiredTime"`
	Status      uint8  `gorm:"not null;index:idx_user_token_status" json:"status"`
}

func NewUserToken() *UserToken {
	return &UserToken{}
}

func (m *UserToken) TableName() string {
	return "user_token"
}

func (m *UserToken) CreateUserToken(data interface{}) *message.CommonResponse {
	resp := message.NewCommonResponse()
	d := data.(form.UserTokenMsgDealReq)

	userToken := &UserToken{
		Token:       d.Token,
		UserId:      d.UserId,
		ExpiredTime: d.ExpireTime,
		Status:      1,
	}

	result := global.MysqlDB.Create(userToken)
	if result.Error != nil {
		resp.Status = message.ModelError
		resp.Message = result.Error.Error()
		return resp
	}

	resp.Message = "userToken创建成功"
	return resp
}

// 从数据库中查询userToke的Status时候，需要对比过期时间
