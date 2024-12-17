package global

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	MysqlDB     *gorm.DB
	ConfigMysql mysql
	Validate    *validator.Validate
)

const (
	DealInfoSuccess = 0
	DealServiceFail = 1
	DealModelFail   = 2
)
