package global

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	// 多个线程使用MysqlDB是否会发生并发问题，gorm.DB中我没找到并发锁的使用。
	MysqlDB     *gorm.DB
	ConfigMysql mysql
	Validate    *validator.Validate
)

const (
	DealInfoSuccess = 0
	DealServiceFail = 1
	DealModelFail   = 2
)
