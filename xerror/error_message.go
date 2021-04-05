package xerror

type ErrorMessage struct {
	ParamsError string
}

var Message = ErrorMessage{ParamsError: "params is invalid"}
