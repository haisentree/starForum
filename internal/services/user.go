package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"math/rand"
	"starForum/internal/global/form"
	"starForum/pkg/password"
	"strings"
	"time"

	"starForum/internal/global"
	"starForum/internal/global/message"
	"starForum/internal/models"
)

var UserService = newUserService()

func newUserService() *userService {
	return &userService{}
}

type userService struct {
}

func (u *userService) SigupInfoDeal(data interface{}) *message.CommonResponse {
	resp := message.NewCommonResponse()
	// 校验验证码
	d := data.(form.SignupInfoMsgReq)
	// 跳过验证码debug
	if global.ConfigServer.Debug != true {
		if global.CaptchaStore.Verify(d.CaptchaId, d.CaptchaAnswer, true) == false {

			resp.Status = message.ServiceError
			resp.Message = "验证码校验错误"
			return resp
		}
	}

	d.Password = password.EncodePassword(d.Password)
	// 判断邮件地址是否被注册过
	userModel := models.NewUser()
	fmt.Println("service:", d.Email)
	respModel := userModel.FindUserByEmail(d.Email)
	//fmt.Println("data:", respModel.Data)
	if respModel.Data != nil {
		resp.Status = message.ServiceError
		resp.Message = "注册失败，该邮件已经被注册"
		return resp

	}
	// 发送验证码
	randCode := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	global.EmailSender.SendSampleCode(randCode, d.Email)
	// 将注册信息以email为key，存储在缓存中
	myCache := form.SignupInfoMsgCache{
		Email:     d.Email,
		Username:  d.Username,
		Password:  d.Password,
		Nickname:  d.Nickname,
		Avatar:    d.Avatar,
		EmailCode: randCode,
	}
	global.Cache.Set(d.Email, myCache, 5*time.Minute)

	resp.Message = "注册信息提交成功,已发送邮件验证码。请尽快验证，有效期5分钟"
	return resp
}

func (u *userService) SignupEmailVerify(data interface{}) *message.CommonResponse {
	resp := message.NewCommonResponse()
	d := data.(form.SigupEmailVerifyMsgReq)
	// 从cache中校验验证码
	item, _ := global.Cache.Get(d.Email)
	userSigupCache := item.(form.SignupInfoMsgCache)
	if userSigupCache.EmailCode != d.EmailCode {
		resp.Status = message.ServiceError
		resp.Message = "邮件验证码错误"
		return resp
	}
	// 将注册信息存储在数据库中
	userModel := models.NewUser()
	userModel.Username = userSigupCache.Username
	userModel.Email = userSigupCache.Email
	userModel.Password = userSigupCache.Password
	userModel.Nickname = userSigupCache.Nickname
	userModel.Avatar = userSigupCache.Avatar

	respModel := models.CreateUser(userModel)
	if respModel.Status != message.SuccessStatus {
		resp.Status = message.ModelError
		resp.Message = "创建用户失败"
		return resp
	}

	resp.Message = "创建用户成功"
	return resp
}
func (u *userService) UserLogin(data interface{}) *message.CommonResponse {
	resp := message.NewCommonResponse()
	d := data.(form.UserLoginMsgReq)
	// 通过查询email从数据库中获取信息
	userModel := models.NewUser()
	respModel := userModel.FindUserByEmail(d.Email)
	if respModel.Status != message.SuccessStatus {
		return respModel
	}
	userData := respModel.Data.(*models.User)
	myUserResp := &form.UserLoginMsgResp{
		UserId:   userData.ID,
		Email:    userData.Email,
		Username: userData.Username,
		Nickname: userData.Nickname,
		Avatar:   userData.Avatar,
		Token:    "",
	}
	// 生成token,并且存储在缓存中，还要存储在数据库中，以token:userID的形式
	respDeal := generateTokenSaveToCacheToDB(myUserResp.UserId)
	if respDeal.Status != message.SuccessStatus {
		resp.Status = message.ServiceError
		return respDeal
	}
	myUserResp.Token = respDeal.Data.(form.UserTokenMsgDealResp).Token

	resp.Message = "token创建成功，用户登录成功"
	resp.Data = myUserResp
	return resp
}

func (u *userService) GetCurrentUserByToken(data interface{}) *message.CommonResponse {
	resp := message.NewCommonResponse()
	d := data.(string)
	// 1.从缓存中查找token(暂时不写2，缓存过期时间，即使登录失效时间)
	myCacheUserID, ok := global.Cache.Get(d)
	if ok == false {
		resp.Status = message.ServiceError
		resp.Message = "token不存在，用户未登录"
		return resp
	}

	userModel := models.NewUser()
	respModel := userModel.FindUserByID(myCacheUserID.(uint))
	if respModel.Status != message.SuccessStatus {
		return respModel
	}
	respData := respModel.Data.(*models.User)
	// 2.缓存中没有，从数据库中查询token，并且存储在缓存中token:
	// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

	myData := &form.CurrentUserMsgResp{
		UserID:   respData.ID,
		Email:    respData.Email,
		Username: respData.Username,
		Nickname: respData.Nickname,
		Avatar:   respData.Avatar,
	}
	resp.Data = myData
	resp.Message = "用户信息查询成功"
	return resp
}

// 用户登录成功用户时候生成token
func generateTokenSaveToCacheToDB(userId uint) *message.CommonResponse {
	u := uuid.New()
	t := time.Now()
	addTime := time.Duration(global.ConfigCache.ExpireTime) * time.Hour
	expireTime := t.Add(addTime)

	userToken := form.UserTokenMsgDealReq{}
	userToken.Token = strings.ReplaceAll(u.String(), "-", "")
	userToken.UserId = userId
	userToken.ExpireTime = expireTime.Unix()

	// 数据库中创建token
	m := models.NewUserToken()
	respModel := m.CreateUserToken(userToken)
	if respModel.Status != message.SuccessStatus {
		return respModel
	}
	// 存储token在缓存中
	global.Cache.Set(userToken.Token, userToken.UserId, cache.DefaultExpiration)

	respModel.Data = form.UserTokenMsgDealResp{
		UserId: userId,
		Token:  userToken.Token,
	}
	return respModel
}

// 处理用户token相关逻辑

// 通过token查询用户信息,收先通过缓存查取，没有去查询数据库并且进行缓存
func getByToken() {

}
