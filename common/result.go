package common

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func Success() Result {
	return Result{
		Code: SUCCESS,
		Msg:  "success",
	}
}

func SuccessWithData(data interface{}) Result {
	return Result{
		Code: SUCCESS,
		Msg:  "success",
		Data: data,
	}
}

func Error(msg string) Result {
	return Result{
		Code: ERROR,
		Msg:  msg,
	}
}

func CustomResult(code int, msg string, data interface{}) Result {
	return Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
