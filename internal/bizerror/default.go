package bizerror

import "errors"

const (
	ErrCodeDefaultBadParams = 1001001
)

var defaultErrCodeMap = ErrCodeMap{
	ErrCodeDefaultBadParams: errors.New("参数有误"),
}
