package errno

var (
	SUCCESS        = &Errno{Code: 0, Message: "OK"}
	ERROR          = &Errno{Code: 500, Message: "Internal server error"}
	INVALID_PARAMS = &Errno{Code: 400, Message: "请求参数错误"}

	ERROR_BIND              = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ERROR_EXIST_TAG         = &Errno{Code: 10002, Message: "已存在该标签名称"}
	ERROR_NOT_EXIST_TAG     = &Errno{Code: 10003, Message: "该标签不存在"}
	ERROR_NOT_EXIST_ARTICLE = &Errno{Code: 10004, Message: "该文章不存在"}

	ERROR_AUTH_CHECK_TOKEN_FAIL    = &Errno{Code: 20001, Message: "Token鉴权失败"}
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = &Errno{Code: 20002, Message: "Token已超时"}
	ERROR_AUTH_TOKEN               = &Errno{Code: 20003, Message: "Token生成失败"}
	ERROR_AUTH                     = &Errno{Code: 20004, Message: "Token错误"}
	EROOR_USER_NOT_FOUND           = &Errno{Code: 20102, Message: "The user was not found."}
)
