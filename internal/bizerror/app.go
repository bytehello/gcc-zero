package bizerror

import "errors"

var (
	ErrCodeAppAddBadParams = 1021001
	ErrCodeAppFind         = 1021002
	ErrCodeAppAddExisted   = 1021003
)
var AppCodeMap = ErrCodeMap{
	ErrCodeAppAddBadParams: errors.New("参数有误"),
	ErrCodeAppFind:         errors.New("查询失败"),
	ErrCodeAppAddExisted:   errors.New("app已经存在"),
}
