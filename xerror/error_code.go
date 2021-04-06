package xerror

type ErrorCode struct {
	CParamsError    string
	CreateUserError string
	SGetLockeErr    string

	SSqlExecuteErr string
	SUnexpectedErr string

	BDataIsNotAllow string
}

var Code = ErrorCode{
	CParamsError:    "100001",
	CreateUserError: "100002",

	SGetLockeErr:   "100003",
	SSqlExecuteErr: "100003",
	SUnexpectedErr: "100004",

	BDataIsNotAllow: "100005",
}
