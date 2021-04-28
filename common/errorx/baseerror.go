package errorx

var defaultErrCode int64 = 1

type CodeError struct {
	Code int64
	Msg  string
}

type CodeErrorResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int64, msg string) *CodeError {
	return &CodeError{
		Code: code,
		Msg:  msg,
	}
}

func DefaultCodeError(msg string) *CodeError {
	return NewCodeError(defaultErrCode, msg)
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
