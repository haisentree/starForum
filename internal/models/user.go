package models

import (
	"starForum/internal/global"
	"starForum/internal/global/form"
	"starForum/internal/global/message"
	"time"
)

type User struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Username  string `gorm:"size:32;unique;" json:"username"` // 用户名
	Email     string `gorm:"size:128;unique;" json:"email"`   // 邮箱
	Password  string `gorm:"size:512;" json:"password"`
	Nickname  string `gorm:"size:16;" json:"nickname"` // 昵称
	Avatar    string `gorm:"type:text" json:"avatar"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser() *User {
	m := &User{}
	return m
}
func (m *User) TableName() string {
	return "t_user"
}

func (m *User) CreateUser(data interface{}) message.CommonDealInfo {
	dealInfo := message.NewCommonDealInfo(nil)

	d := data.(form.SignupMsgReq)
	m.Username = d.Username
	m.Email = d.Email
	m.Password = d.Password
	m.Nickname = d.Nickname
	m.Avatar = d.Avatar

	result := global.MysqlDB.Create(m)
	if result.Error != nil {
		dealInfo.Error = 1
		dealInfo.Message = result.Error.Error()
		return dealInfo
	}
	return dealInfo
}

// ========================================UserToken========================================
type UserToken struct {
	BaseModel
	Token     string `gorm:"size:32;unique;not null" json:"token"`
	UserId    int64  `gorm:"not null;index:idx_user_token_user_id;" json:"userId"`
	ExpiredAt int64  `gorm:"not null" json:"expiredAt"`
	Status    int    `gorm:"type:int(11);not null;index:idx_user_token_status" json:"status"`
}

func NewUserToken() *UserToken {
	m := &UserToken{}
	return m
}

func (m *UserToken) TableName() string {
	return "t_user_token"
}

func (m *UserToken) CreateUsertoken(data interface{}) message.CommonDealInfo {
	dealInfo := message.NewCommonDealInfo(nil)
	userToken := &UserToken{
		Token: m.Token,
	}

}
