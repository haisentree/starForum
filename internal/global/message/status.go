package message

// ===============================StatusCode====================================
const (
	SuccessStatus = 200
	SucessMessage = "success"
)

// 发生在controller中处理数据错误
const (
	ControllerError        = 1000
	ControllerErrorMessage = "controller error"
)

// 发生在service中处理数据错误
const (
	ServiceError        = 1000
	ServiceErrorMessage = "service error"
)

// 发生在model中处理数据错误
const (
	ModelError        = 2000
	ModelErrorMessage = "model error"
	ModelFindNone     = 2001
)
