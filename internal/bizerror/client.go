package bizerror

import "errors"

const (
	ErrCodeClientFind           = 1041001
	ErrCodeClientRecordNotFound = 1041002
)

var ClientMap = ErrCodeMap{
	ErrCodeClientFind:           errors.New("查询失败"),
	ErrCodeClientRecordNotFound: errors.New("没有结果"),
}
