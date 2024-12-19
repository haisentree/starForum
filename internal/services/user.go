package services

import (
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"starForum/internal/global/cacheInfo"
	"starForum/internal/global/form"
	"starForum/pkg/password"
	"time"

	"starForum/internal/global"
	"starForum/internal/global/message"
	"starForum/internal/models"
	"strings"
)

var UserService = newUserService()

func newUserService() *userService {
	return &userService{}
}

type userService struct {
}

func SigupInfoDeal(data interface{}) message.CommonDealInfo {
	res := message.NewCommonDealInfo(nil)

	d := data.(form.SignupInfoMsgReq)
	if global.CaptchaStore.Verify(d.CaptchaId, d.CaptchaAnswer, true) == false {
		res.Error = global.DealServiceFail
		res.Message = "验证码校验错误"
		return res
	}
	d.Password = password.EncodePassword(d.Password)
	// 判断邮件地址是否被注册过

	// 发送验证码

	// 存储再存储再缓存中

	res.Message = "注册信息提交成功,已发送邮件验证码，请尽快验证"
	return res
}

func CreateUser(data interface{}) message.CommonDealInfo {

	u := models.NewUser()
	dealInfo := u.CreateUser(data.(form.LoginMsgReq))
	if dealInfo.Error != global.DealInfoSuccess {
		return dealInfo
	}
	dealInfo.Data = u
	return dealInfo
}

// 处理用户token相关逻辑

// 用户登录和注册成功用户时候生成token
func generateToken(data interface{}) message.CommonDealInfo {
	u := uuid.New()
	strings.ReplaceAll(u.String(), "-", "")
	t := time.Now()
	addTime := time.Duration(global.ConfigCache.ExpireTime) * time.Hour
	expireTime := t.Add(addTime)

	userToken := form.UserTokenMsgDeal{}
	userToken.Token = u.String()
	userToken.UserId = data.(form.UserTokenMsgReq).UserId
	userToken.ExpireTime = expireTime.Unix()

	m := models.NewUserToken()
	dealInfo := m.CreateUsertoken(userToken)
	if dealInfo.Error != global.DealInfoSuccess {
		dealInfo.Message = "service创建token失败"
		return dealInfo
	}
	// 存储token在缓存中
	global.Cache.Set("foo", cacheInfo.UserCache{
		UserName: "test",
	}, cache.DefaultExpiration)

	dealInfo.Data = userToken
	return dealInfo
}

// 通过token查询用户信息,收先通过缓存查取，没有去查询数据库并且进行缓存
func getByToken() {

}
