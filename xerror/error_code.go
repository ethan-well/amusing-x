package xerror

type ErrorCode struct {
	ParamsError string
}

var Code = ErrorCode{ParamsError: "100001"}
