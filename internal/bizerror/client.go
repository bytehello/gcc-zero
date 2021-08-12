package bizerror

import "errors"

const (
	ErrCodeClientFind = 1041001
)

var ClientMap = ErrCodeMap{
	ErrCodeClientFind: errors.New("查询失败"),
}
