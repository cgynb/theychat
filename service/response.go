package service

type Resp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func RespOk(data any) *Resp {
	return &Resp{
		Code:    200,
		Message: "ok",
		Data:    data,
	}
}

func RespErr(code int32, message string) *Resp {
	return &Resp{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
