package server

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/mojocn/base64Captcha"
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"starForum/internal/global"
	"starForum/internal/global/config"
	"starForum/internal/models"
	myemail "starForum/pkg/email"
	"time"
)

func init() {
	initConfig()
	initDB()
	initValidate()
	initCache()
	initCaptcha()
	initEmail()
}

func initConfig() {
	viper.SetConfigName("star-forum.dev.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.ReadInConfig()
	var config config.Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Sprintf("unable to decode config, %v", err))
	}
	global.ConfigMysql = config.Mysql
	global.ConfigCache = config.Cache
	global.ConfigEmail = config.Email
	global.ConfigServer = config.Server
}

func initDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		global.ConfigMysql.Username,
		global.ConfigMysql.Password,
		global.ConfigMysql.Host,
		global.ConfigMysql.Port,
		global.ConfigMysql.DBName,
		global.ConfigMysql.Charset,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true, // 单数形式User{}，创建t_user
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(models.Models...)
	global.MysqlDB = db
}

func initValidate() {
	v := validator.New()
	global.Validate = v
}

func initCache() {
	expireTime := time.Duration(global.ConfigCache.CacheTime)
	cleanTime := time.Duration(global.ConfigCache.CleanTime)
	c := cache.New(expireTime*time.Minute, cleanTime*time.Minute)
	global.Cache = c
}

func initCaptcha() {
	digitDriver := base64Captcha.DriverDigit{
		Height:   50,
		Width:    200,
		Length:   4,   //验证码长度
		MaxSkew:  0.7, //倾斜
		DotCount: 1,   //背景的点数，越大，字体越模糊
	}
	global.CaptchaGenerate = base64Captcha.NewCaptcha(&digitDriver, global.CaptchaStore)
}

func initEmail() {
	Subject := "验证码发送"
	e := myemail.NewEmailSender(
		global.ConfigEmail.Sender,
		Subject,
		global.ConfigEmail.Username,
		global.ConfigEmail.Password,
		global.ConfigEmail.Host,
		global.ConfigEmail.Port,
		global.ConfigEmail.TLS,
	)

	global.EmailSender = e
}
