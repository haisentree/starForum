package models

import (
	"gorm.io/gorm"
	"starForum/internal/global"
	"starForum/internal/global/form"
	"starForum/internal/global/message"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:32;unique;" json:"username" form:"username"` // 用户名
	Email    string `gorm:"size:128;unique;" json:"email" form:"email"`      // 邮箱
	Password string `gorm:"size:512;" json:"password" form:"password"`
	Nickname string `gorm:"size:16;" json:"nickname" form:"nickname"` // 昵称
	Avatar   string `gorm:"type:text" json:"avatar" form:"avatar"`
}

func (User) TableName() string {
	return "t_user"
}

func NewUser() *User {
	c := &User{}
	return c
}

func (c *User) CreateUser(data interface{}) message.CommonDealInfo {
	dealInfo := message.NewCommonDealInfo(nil)

	d := data.(form.SignupMsgReq)
	c.Username = d.Username
	c.Email = d.Email
	c.Password = d.Password
	c.Nickname = d.Nickname
	c.Avatar = d.Avatar

	result := global.MysqlDB.Create(c)
	if result.Error != nil {
		dealInfo.Error = 1
		dealInfo.Message = result.Error.Error()
		return dealInfo
	}
	return dealInfo
}
