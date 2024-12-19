package global

import (
	"github.com/go-playground/validator/v10"
	"github.com/mojocn/base64Captcha"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"starForum/internal/global/config"
)

var (
	// 多个线程使用MysqlDB是否会发生并发问题，gorm.DB中我没找到并发锁的使用。
	MysqlDB         *gorm.DB
	Validate        *validator.Validate
	Cache           *cache.Cache
	CaptchaGenerate *base64Captcha.Captcha
	CaptchaStore    = base64Captcha.DefaultMemStore
)

var (
	ConfigMysql config.Mysql
	ConfigCache config.Cache
)

const (
	DealInfoSuccess = 0
	DealServiceFail = 1
	DealModelFail   = 2
)
