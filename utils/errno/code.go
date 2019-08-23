package errno

var (
	// Common errors
	OK                  = &Errno{Code: 1, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "内部服务错误。"}
	ErrBind             = &Errno{Code: 10002, Message: "请求体绑定到数据结构错误。"}

	ErrValidation    = &Errno{Code: 20001, Message: "参数验证失败。"}
	ErrDatabase      = &Errno{Code: 20002, Message: "数据库连接失败。"}
	ErrToken         = &Errno{Code: 20003, Message: "请求Token验证失败。"}
	ErrSign          = &Errno{Code: 20004, Message: "请求签名验证失败。"}
	ErrMissingHeader = &Errno{Code: 20005, Message: "请求头中没有Authorization字段。"}

	ErrUserNotFound      = &Errno{Code: 20102, Message: "用户不存在。"}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "用户密码输入错误"}
)
