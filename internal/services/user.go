package services

import (
	"starForum/internal/global"
	"starForum/internal/global/message"
	"starForum/internal/models"
)

func CreateUser(data interface{}) message.CommonDealInfo {

	u := models.NewUser()
	dealInfo := u.CreateUser(data)
	if dealInfo.Error != global.DealInfoSuccess {
		return dealInfo
	}
	dealInfo.Data = u
	return dealInfo
}
