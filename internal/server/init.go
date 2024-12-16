package server

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"starForum/internal/global"
)

func init() {
	initConfig()
	initDB()
}

func initConfig() {
	viper.SetConfigName("star-forum.dev.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.ReadInConfig()
	var config global.Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Sprintf("unable to decode config, %v", err))
	}
	global.ConfigMysql = config.Mysql
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	global.MysqlDB = db
}
