package services

import (
	"github.com/google/uuid"
	"starForum/internal/global/form"
	"starForum/internal/server"

	"starForum/internal/global"
	"starForum/internal/global/message"
	"starForum/internal/models"
	"strings"
)

func CreateUser(data interface{}) message.CommonDealInfo {

	u := models.NewUser()
	dealInfo := u.CreateUser(data.(form.SignupMsgReq))
	if dealInfo.Error != global.DealInfoSuccess {
		return dealInfo
	}
	dealInfo.Data = u
	return dealInfo
}

// 处理用户token相关逻辑

func generate(data interface{}) message.CommonDealInfo {
	u := uuid.New()
	strings.ReplaceAll(u.String(), "-", "")

	userToken := form.UserTokenMsgDeal{}
	userToken.Token = u.String()
	userToken.UserId = data.(form.UserTokenMsgReq).UserId

	m := models.NewUserToken()
	dealInfo := m.CreateUsertoken(form.UserTokenMsgDeal{})
	if dealInfo.Error != global.DealInfoSuccess {
		return dealInfo
	}

	dealInfo.Data = userToken
	return dealInfo

}
