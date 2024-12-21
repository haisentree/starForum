package message

import (
	"github.com/kataras/iris/v12"
)

// =====================后端响应信息给前端的数据================================
type CommonResponse struct {
	Status  int32
	Message string
	Data    interface{}
}

// models -> service    service -> controller
func NewCommonResponse() *CommonResponse {
	return &CommonResponse{
		Status:  SuccessStatus,
		Message: SucessMessage,
		Data:    nil,
	}
}

// controller -> web
func (c *CommonResponse) JsonCommonResponse() iris.Map {
	return iris.Map{
		"status":  c.Status,
		"message": c.Message,
		"data":    c.Data,
	}
}

// 用于传递controllers与service与model三者之间的错误信息
// 约定:0表示没有错误,1表示model处理错误,2表示service处理错误

type CommonDealInfo struct {
	Error   uint8
	Message string
	Data    interface{}
}

func NewCommonDealInfo(d interface{}) CommonDealInfo {
	return CommonDealInfo{
		Error:   0,
		Message: "success",
		Data:    d,
	}
}
