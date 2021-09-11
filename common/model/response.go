package model

const (
	Success   = 0  // 成功
	Unknown   = -1 // 未知错误，一般性错误
	WrongArgs = -2 // 收到错误的参数
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Tag  string      `json:"tag,omitempty"`
	Data interface{} `json:"data"`
}

func NewDataResponse(data interface{}, tag string) *Response {
	var r Response
	r.Msg = "success"
	r.Code = Success
	r.Data = data
	r.Tag = tag
	return &r
}

func SimpleResponse(code int, msg string) *Response {
	var r Response
	r.Msg = msg
	r.Code = code
	return &r
}

func NewBindFailedResponse(tag string) *Response {
	return &Response{Code: WrongArgs, Msg: "wrong argument", Tag: tag}
}
