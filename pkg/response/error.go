package response

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

type CustomErrors struct {
	BusinessError CustomError
	ValidateError CustomError
}

var Errors = CustomErrors{
	BusinessError: CustomError{10000, "业务错误"},
	ValidateError: CustomError{20000, "请求参数错误"},
}
