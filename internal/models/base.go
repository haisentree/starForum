package models

var Models = []interface{}{
	&User{}, &UserToken{},
}

type BaseModel struct {
	Id uint64 `gorm:"primarykey"`
}
