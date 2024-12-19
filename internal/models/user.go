package models

import (
	"gorm.io/gorm"
	"starForum/internal/global"
	"starForum/internal/global/form"
	"starForum/internal/global/message"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:32;unique;" json:"username"` // 用户名
	Email    string `gorm:"size:128;unique;" json:"email"`   // 邮箱
	Password string `gorm:"size:512;" json:"password"`
	Nickname string `gorm:"size:16;" json:"nickname"` // 昵称
	Avatar   string `gorm:"type:text" json:"avatar"`
}

func NewUser() *User {
	m := &User{}
	return m
}
func (m *User) TableName() string {
	return "user"
}

func (m *User) CreateUser(data interface{}) message.CommonDealInfo {
	dealInfo := message.NewCommonDealInfo(nil)

	d := data.(form.LoginMsgReq)
	m.Username = d.Username
	m.Email = d.Email
	m.Password = d.Password
	m.Nickname = d.Nickname
	m.Avatar = d.Avatar

	result := global.MysqlDB.Create(m)
	if result.Error != nil {
		dealInfo.Error = global.DealModelFail
		dealInfo.Message = result.Error.Error()
		return dealInfo
	}
	return dealInfo
}

// ========================================UserToken========================================
type UserToken struct {
	gorm.Model
	Token       string `gorm:"size:32;unique;not null" json:"token"`
	UserId      uint   `gorm:"not null;index:idx_user_token_user_id;" json:"userId"`
	ExpiredTime int64  `gorm:"not null" json:"expiredTime"`
	Status      uint8  `gorm:"not null;index:idx_user_token_status" json:"status"`
}

func NewUserToken() *UserToken {
	m := &UserToken{}
	return m
}

func (m *UserToken) TableName() string {
	return "user_token"
}

func (m *UserToken) CreateUsertoken(data interface{}) message.CommonDealInfo {
	dealInfo := message.NewCommonDealInfo(nil)
	d := data.(form.UserTokenMsgDeal)
	userToken := &UserToken{
		Token:       d.Token,
		UserId:      d.UserId,
		ExpiredTime: d.ExpireTime,
		Status:      0,
	}

	result := global.MysqlDB.Create(userToken)
	if result.Error != nil {
		dealInfo.Error = global.DealModelFail
		dealInfo.Message = result.Error.Error()
		return dealInfo
	}

	return dealInfo
}

// 从数据库中查询userToke的Status时候，需要对比过期时间
