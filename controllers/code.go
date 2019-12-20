package controllers

var (
	// Common errors
	SUCCESS             = &Errno{Code: 200, Message: "成功"}
	InternalServerError = &Errno{Code: 10001, Message: "内部服务错误"}
	ErrBind             = &Errno{Code: 10002, Message: "参数错误"}

	ErrDatabase     = &Errno{Code: 20001, Message: "数据库错误"}
	ErrToken        = &Errno{Code: 20002, Message: "签发令牌出错"}
	ErrNoPermission = &Errno{Code: 401, Message: "无权限"}

	// user errors
	ErrUserNotFound = &Errno{Code: 20101, Message: "用户未注册"}
	ErrUserExist    = &Errno{Code: 20102, Message: "用户已存在"}
)
