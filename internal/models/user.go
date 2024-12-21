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
	return &User{}
}
func (m *User) TableName() string {
	return "user"
}

func CreateUser2(data *User) *message.CommonResponse {
	resp := message.NewCommonResponse()
	result := global.MysqlDB.Create(data)
	if result.Error != nil {
		resp.Status = message.ModelError
		resp.Message = result.Error.Error()
		return resp
	}
	return resp
}

func (m *User) FindUserByEmail(email string) *message.CommonResponse {
	respModel := message.NewCommonResponse()
	user := &User{
		Email: email,
	}
	result := global.MysqlDB.First(&user)
	if result.RowsAffected == 0 {
		respModel.Status = message.ModelFindNone
		respModel.Message = "数据不存在数据库中"
		return respModel
	}
	respModel.Data = user
	return respModel
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
